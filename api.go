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
	HTTP_PORT = 8080
)

//Article xxxx
type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt int    `json:"created_at"`
}

type JSON struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/detail", Detail)

	fmt.Println(fmt.Sprintf("listen %s:%d", HTTP_HOST, HTTP_PORT))
	http.ListenAndServe(fmt.Sprintf("%s:%d", HTTP_HOST, HTTP_PORT), nil)
}

//Detail 详情
func Detail(w http.ResponseWriter, r *http.Request) {
	data := &JSON{}
	data.Status = false
	data.Message = "failure"

	port := r.FormValue("port")
	url := fmt.Sprintf("http://127.0.0.1:22999/api/refresh_sessions/%s", port)
	//url = "http://baidu.com"

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
	return
}

//JSONReturn 输出
func JSONReturn(w http.ResponseWriter, content JSON) {
	detailJSON, _ := json.Marshal(content)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(detailJSON)
	return
}
