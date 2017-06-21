package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/farmerx/elasticsql"
	"github.com/mattbaird/elastigo/lib"
)

// InitOptions struct
type InitOptions struct {
	Mux *http.ServeMux
}

// Search struct
type Search struct {
	mux *http.ServeMux
}

// NewSearch init
func NewSearch(options InitOptions) *Search {
	se := new(Search)
	se.mux = options.Mux
	se.RegHandler()
	return se
}

// RegHandler 路由注册机制
func (se *Search) RegHandler() {
	se.mux.HandleFunc("/_ping", se.ping)
	se.mux.HandleFunc("/_cat/health", se.health)
	se.mux.HandleFunc("/_cat/indices", se.indices)
	se.mux.HandleFunc("/_cluster/stats", se.clusterStats)
	se.mux.HandleFunc("/_template", se.template)
	se.mux.HandleFunc("/_search", se.search)
	se.mux.HandleFunc("/_upload/template", se.uploadTemplate)
	se.mux.HandleFunc("/_template/handle", se.handleTemplate)
	se.mux.HandleFunc("/_sql2dsl", se.sqlConvert)
	se.mux.HandleFunc("/_indices/delete", se.indexDelete)
}

func (se *Search) sqlConvert(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"error": "服务端错误：参数解析失败。"}))
		return
	}
	esql := elasticsql.NewElasticSQL(elasticsql.InitOptions{})
	_, dsl, err := esql.SQLConvert(param.SQL)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"error": err.Error()}))
		return
	}
	w.Write(jsonEncode(0, []interface{}{dsl}))
}

// {
// result: 0,
// data: {
// cluster_name: "es",
// name: "es.10.187.102.8.0",
// tagline: "You Know, for Search",
// version: {
// build_hash: "d7d6a2ae333a09006b54954759dd41dbe9aefb82",
// build_snapshot: false,
// build_timestamp: "2016-06-02T09:13:39Z",
// lucene_version: "5.2.1",
// number: "2.0.0"
// }
// }
// }
func (se *Search) ping(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	body, err := curlGet(param.Serverhost, 10*time.Second)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	var dBody interface{}
	if err := json.Unmarshal(body, &dBody); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	w.Write(jsonEncode(0, dBody))
}

// Health 集群健康状态
// {
//   "result": 0,
//   "data": {
//     "active_primary_shards": 665,
//     "active_shards": 665,
//     "active_shards_percent_as_number": 100,
//     "cluster_name": "es",
//     "delayed_unassigned_shards": 0,
//     "initializing_shards": 0,
//     "number_of_data_nodes": 1,
//     "number_of_in_flight_fetch": 0,
//     "number_of_nodes": 1,
//     "number_of_pending_tasks": 0,
//     "relocating_shards": 0,
//     "status": "green",
//     "task_max_waiting_in_queue_millis": 0,
//     "timed_out": false,
//     "unassigned_shards": 0
//   }
func (se *Search) health(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	url := fmt.Sprintf(`%s/_cluster/health`, param.Serverhost)
	body, err := curlGet(url, 60*time.Second)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	var dBody interface{}
	if err := json.Unmarshal(body, &dBody); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	w.Write(jsonEncode(0, dBody))
}

func (se *Search) clusterStats(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	url := fmt.Sprintf(`%s/_cluster/stats?human&pretty`, param.Serverhost)
	body, err := curlGet(url, 60*time.Second)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	var stats ClusterStats
	if err := json.Unmarshal(body, &stats); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	var elasticStats = Stats{}
	elasticStats.ClusterName = stats.ClusterName
	elasticStats.Docs = stats.Indices.Docs.Count
	elasticStats.FieldDataMemUsed = stats.Indices.Fielddata.MemorySizeInBytes
	elasticStats.FsFree = stats.Nodes.Fs.FreeInBytes
	elasticStats.FsTotal = stats.Nodes.Fs.TotalInBytes
	elasticStats.Indices = stats.Indices.Count
	elasticStats.JvmMaxUptime = stats.Nodes.Jvm.MaxUptime
	elasticStats.JvmMemMaxUse = stats.Nodes.Jvm.Mem.HeapMaxInBytes
	elasticStats.JvmMemUsed = stats.Nodes.Jvm.Mem.HeapUsedInBytes
	elasticStats.JvmThreads = stats.Nodes.Jvm.Threads
	for _, item := range stats.Nodes.Jvm.Versions {
		elasticStats.JvmVersions += item.Version + ","
	}
	elasticStats.MemFree = stats.Nodes.Os.Mem.FreeInBytes
	elasticStats.MemToatl = stats.Nodes.Os.Mem.TotalInBytes
	elasticStats.MemUsed = stats.Nodes.Os.Mem.UsedInBytes
	elasticStats.Plugins = stats.Nodes.Plugins
	elasticStats.QueryCache = stats.Indices.QueryCache
	elasticStats.Shards = stats.Indices.Shards.Total
	elasticStats.Size = stats.Indices.Store.Size
	elasticStats.Status = stats.Status
	for _, item := range stats.Nodes.Os.Names {
		elasticStats.Systems += item.Name + ","
	}
	elasticStats.Versions = stats.Nodes.Versions
	elasticStats.Timestamp = stats.Timestamp
	w.Write(jsonEncode(0, elasticStats))
}

func (se *Search) handleTemplate(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	u, err := url.Parse(param.Serverhost)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	c := elastigo.NewConn()
  c.Username = u.User.Username()
  c.Password, _ = u.User.Password()
	c.Domain = strings.Split(u.Host, ":")[0]
	c.Port = strings.Split(u.Host, ":")[1]
	c.DecayDuration = 0
	url := fmt.Sprintf(`/_template/%s`, param.Indices)
	result, err := c.DoCommand(param.Method, url, nil, nil)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	var dBody interface{}
	if err := json.Unmarshal(result, &dBody); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	w.Write(jsonEncode(0, dBody))
}

//TemplateContent ...
func (se *Search) template(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	url := fmt.Sprintf(`%s/_template`, param.Serverhost)
	body, err := curlGet(url, 600*time.Second)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	var dBody map[string]struct {
		Template string      `json:"template"`
		Settings interface{} `json:"settings"`
		Mappings map[string]struct {
			Dynamic    bool                   `json:"dynamic"`
			Properties map[string]interface{} `json:"properties"`
		}
		Aliases interface{} `json:"aliases"`
	}
	if err := json.Unmarshal(body, &dBody); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	var templateName = make([]string, 0)
	var typeCount = make([]int, 0)
	var tempreflect = make([]interface{}, 0)
	for key, value := range dBody {
		templateName = append(templateName, key)
		typeCount = append(typeCount, len(value.Mappings))
		var tmp = make([]string, 0)
		for name := range value.Mappings {
			tmp = append(tmp, `【`+name+`】`)
		}
		tempreflect = append(tempreflect, map[string]interface{}{
			`name`:    key,
			`type`:    strings.Join(tmp, ``),
			`visible`: false,
		})
	}
	var result = map[string]interface{}{
		`count`:        len(dBody),
		`typeCount`:    typeCount,
		`templateName`: templateName,
		`tmpreflect`:   tempreflect,
	}
	w.Write(jsonEncode(0, result))
	return
}

//Indices ...
func (se *Search) indices(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	url := fmt.Sprintf(`%s/_cat/indices?format=json`, param.Serverhost)
	body, err := curlGet(url, 600*time.Second)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	var indices indicesSort
	if err := json.Unmarshal(body, &indices); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{}))
		return
	}
	sort.Stable(indices)
	var indexes = make([]map[string]interface{}, 0)
	for _, value := range indices {
		indexes = append(indexes, map[string]interface{}{
			`docs`:    value.Count,
			`delete`:  value.Delete,
			`health`:  value.Health,
			`index`:   value.Index,
			`size`:    value.Size,
			`status`:  value.Status,
			`uuid`:    value.UUID,
			`pri`:     value.Pri,
			`prisize`: value.PriSize,
			`rep`:     value.Rep,
			`visible`: false,
		})
	}
	w.Write(jsonEncode(0, indexes))
}

// Search ...
func (se *Search) search(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	u, err := url.Parse(param.Serverhost)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	c := elastigo.NewConn()
  c.Username = u.User.Username()
  c.Password, _ = u.User.Password()
	c.Domain = strings.Split(u.Host, ":")[0]
	c.Port = strings.Split(u.Host, ":")[1]
	c.DecayDuration = 0

	url := fmt.Sprintf(`/%s/%s/%s`, strings.TrimSpace(param.Indices), strings.TrimSpace(param.Types), strings.TrimSpace(param.Options))
	if param.Method == `GET` {
		param.Body = ``
	}
	result, err := c.DoCommand(param.Method, url, nil, param.Body)
	if err != nil {
		w.Write(jsonEncode(0, map[string]interface{}{"info": err.Error()}))
		return
	}
	var dBody interface{}
	if err := json.Unmarshal(result, &dBody); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	w.Write(jsonEncode(0, dBody))
}

// Sizer 文件大小
type Sizer interface {
	Size() int64
}

func (se *Search) uploadTemplate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	if fileSizer, ok := file.(Sizer); ok {
		fileSize := fileSizer.Size() / (1024 * 1024)
		if fileSize > 2 {
			msg := "获取上传文件错误:文件大小超出2M"
			http.Error(w, msg, 500)
			return
		}
	}
	filename := strings.Split(handler.Filename, ".json")[0]
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	u, err := url.Parse(r.FormValue(`serverhost`))
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	c := elastigo.NewConn()
  c.Username = u.User.Username()
  c.Password, _ = u.User.Password()
	c.Domain = strings.Split(u.Host, ":")[0]
	c.Port = strings.Split(u.Host, ":")[1]
	c.DecayDuration = 0
	url := fmt.Sprintf(`/_template/%s`, filename)
	if _, err := c.DoCommand(`PUT`, url, nil, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

type indicesSort []Indices

func (s indicesSort) Len() int      { return len(s) }
func (s indicesSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s indicesSort) Less(i, j int) bool {
	ix := strings.Split(s[i].Index, `-`)
	jx := strings.Split(s[j].Index, `-`)

	if len(ix) <= 1 {
		return false
	}
	if len(jx) <= 1 {
		return true
	}
	return ix[len(ix)-1] > jx[len(jx)-1]
}

// delete Index ...
func (se *Search) indexDelete(w http.ResponseWriter, r *http.Request) {
	param, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	u, err := url.Parse(param.Serverhost)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	c := elastigo.NewConn()
  c.Username = u.User.Username()
  c.Password, _ = u.User.Password()
	c.Domain = strings.Split(u.Host, ":")[0]
	c.Port = strings.Split(u.Host, ":")[1]
	c.DecayDuration = 0
	url := fmt.Sprintf(`/%s`, strings.TrimSpace(param.Index))
	result, err := c.DoCommand(`DELETE`, url, nil, nil)
	if err != nil {
		w.Write(jsonEncode(0, map[string]interface{}{"info": err.Error()}))
		return
	}
	var dBody interface{}
	if err := json.Unmarshal(result, &dBody); err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"info": err.Error()}))
		return
	}
	w.Write(jsonEncode(0, dBody))
}
