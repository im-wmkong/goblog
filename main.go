package main

import (
	"fmt"
	"net/http"
)

func handlerFuc(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "请求路径为："+r.URL.Path)
	} else {
		fmt.Fprint(w, "页面未找到")
	}
}

func main()  {
	http.HandleFunc("/", handlerFuc)
	http.ListenAndServe(":3000", nil)
}
