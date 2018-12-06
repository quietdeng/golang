package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	HTTP_HOST = "0.0.0.0"
	HTTP_PORT = 80
)

type JSON struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/index.php", DetailHandle)
	http.HandleFunc("/recent_stats.php", RecentHandle)

	fmt.Println(fmt.Sprintf("listen %s:%d", HTTP_HOST, HTTP_PORT))
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", HTTP_HOST, HTTP_PORT), nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("启动失败，检查%d端口是否被占用。", HTTP_PORT))
	}
}

//DetailHandle 详情
func DetailHandle(w http.ResponseWriter, r *http.Request) {
	data := &JSON{}
	data.Status = false
	data.Message = "failure"

	port := r.FormValue("port")
	url := fmt.Sprintf("http://127.0.0.1:22999/api/refresh_sessions/%s", port)
	//url = "http://www.baidu.com"

	resp, err := http.Post(
		url,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println("request failure")
		data.Message = "request failure"
		JSONReturn(w, *data)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request ReadAll failure")
		data.Message = "request ReadAll failure" + string(body)
		JSONReturn(w, *data)
		return
	}

	data.Status = true
	data.Message = "success"
	JSONReturn(w, *data)

	fmt.Println("request success")
	fmt.Println("request success: http://127.0.0.1:22999/api/refresh_sessions")
	return
}

func RecentHandle(w http.ResponseWriter, r *http.Request) {
	data := &JSON{}
	data.Status = false
	data.Message = "failure"

	url := "http://127.0.0.1:22999/api/recent_stats"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("request failure")
		data.Message = "request failure"
		JSONReturn(w, *data)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request ReadAll failure")
		data.Message = "request ReadAll failure" + string(body)
		JSONReturn(w, *data)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(body)

	fmt.Println("request success: http://127.0.0.1:22999/api/recent_stats")
	return
}

//JSONReturn 输出
func JSONReturn(w http.ResponseWriter, content JSON) {
	detailJSON, _ := json.Marshal(content)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(detailJSON)
	return
}
