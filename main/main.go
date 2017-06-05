package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"

	"360.cn/elasticHD/main/search"
	_ "360.cn/elasticHD/main/statik"
	"github.com/rakyll/statik/fs"
)

var (
	Stderr io.Writer = os.Stderr // Stderr is the io.Writer to which executed commands write standard error.
	Stdout io.Writer = os.Stdout // Stdout is the io.Writer to which executed commands write standard output.
)

// Config ...
type Config struct {
	ServerHost string
}

var config *Config

func main() {

	parseFlags()
	statikFS, err := fs.New() // 静态文件编译成二进制
	if err != nil {
		log.Fatalf(err.Error())
	}
	http.FileServer(statikFS) // 加载http server file
	var mux = http.NewServeMux()
	// router
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	// 加载search package
	search.NewSearch(search.InitOptions{
		Mux: mux,
	})
	// start http server
	go func(serverAddr string, m *http.ServeMux) {
		if err = http.ListenAndServe(serverAddr, m); err != nil {
			fmt.Println("Cant start server:", err)
			os.Exit(1)
		}
	}(config.ServerHost, mux)
	openPage()
	// 捕捉ctrl+C 退出信号
	handleSignals()
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func openPage() {
	url := fmt.Sprintf("http://%v", config.ServerHost)
	fmt.Println("To view elasticHD console open", url, "in browser")
	var err error
	switch runtime.GOOS {
	case "linux":
		err = runCmd("xdg-open", url)
	case "darwin":
		err = runCmd("open", url)
	case "windows":
		r := strings.NewReplacer("&", "^&")
		err = runCmd("cmd", "/c", "start", r.Replace(url))
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println(err)
	}
}

// runCmd run command opens a new browser window pointing to url.
func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	return cmd.Run()
}

// parseFlags parse flags of program.
func parseFlags() {
	config = &Config{}
	flag.StringVar(&config.ServerHost, "p", "0.0.0.0:9800", "local server address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}
