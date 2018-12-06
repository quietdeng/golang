package route

import (
	"book/controller"
	"net/http"
	"strings"
)

//Rule 路由规则
type Rule struct {
	Method   string
	Pattern  string
	Function func(w http.ResponseWriter, r *http.Request)
}

//Routes 路由列表集合
var Routes []Rule

//BookController 控制器
var BookController = controller.BookController{}

//初始化路由列表
func init() {
	Routes = append(Routes, Rule{"GET", "/list", BookController.Lists})
	Routes = append(Routes, Rule{"POST", "/create", BookController.Create})
	Routes = append(Routes, Rule{"POST", "/update", BookController.Update})
	Routes = append(Routes, Rule{"POST", "/delete", BookController.Delete})
}

type RouteHandle struct{}

//路由
func (*RouteHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Path
	method := strings.ToLower(r.Method)

	//匹配路由
	for _, rule := range Routes {
		if rule.Pattern == URL && strings.ToLower(rule.Method) == method {
			rule.Function(w, r)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Write([]byte("route not match"))
	return
}
