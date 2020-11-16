package main

import (
	"fmt"
	"net/http"
)

func handlerFuc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "h1>Hello, 这里是 goblog</h1>");
}

func main()  {
	http.HandleFunc("/", handlerFuc)
	http.ListenAndServe(":3000", nil)
}
