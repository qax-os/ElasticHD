package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"360.cn/elasticHD/main/search"
	_ "360.cn/elasticHD/main/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
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
	}(`127.0.0.1:9800`, mux)
	// 捕捉ctrl+C 退出信号
	handleSignals()
}

func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
