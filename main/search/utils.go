package search

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// 获取前端传递过来的参数
func getParams(r *http.Request) (param Params, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return param, err
	}
	if err = json.Unmarshal(body, &param); err != nil {
		return param, err
	}
	if param.Serverhost == `` {
		return param, errors.New(`host and port cannot be null`)
	}
	return param, nil
}

// curlGet curl get request
func curlGet(url string, timeout time.Duration) ([]byte, error) {
	client := &http.Client{}
	client.Timeout = timeout
	resp, err := client.Get(url)
	if err != nil {
		return []byte(""), err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	resp.Body.Close()
	return body, nil
}

// json encode result
func jsonEncode(code int, data interface{}) []byte {
	tr := new(ReturnResult)
	tr.Result, tr.Data = code, data
	result, _ := json.Marshal(tr)
	return result
}
