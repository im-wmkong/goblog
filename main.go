package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"
)

var router = mux.NewRouter()
var db *sql.DB

type ArticlesFormData struct {
	Title, Body string
	URL *url.URL
	Errors map[string]string
}

func initDB()  {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}
	// 准备数据库连接池
	db, err = sql.Open("mysql", config.FormatDSN())
	checkError(err)

	// 设置最大连接数
	db.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个连接的过期时间
	db.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接
	err = db.Ping()
	checkError(err)
}

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
		"<p>如有疑惑，请联系我们。</p>")
}

func aritlcesShowHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章 ID：" + id)
}

func aritlcesIndexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"访问文章列表")
}

func aritlcesCreateHandler(w http.ResponseWriter, r *http.Request)  {
	storeURL, _ := router.Get("articles.store").URL()
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, data)
}

func aritlcesStoreHandler(w http.ResponseWriter, r *http.Request)  {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	
	errors := make(map[string]string)
	
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度介于 3～ 40"
	}
	
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需要大于等于 10 个字节"
	}

	if len(errors) == 0 {
		fmt.Fprint(w, "验证通过!<br>")
		fmt.Fprintf(w, "title 的值为: %v <br>", title)
		fmt.Fprintf(w, "body 的值为: %v <br>", body)
	} else {
		storeURL, _ := router.Get("articles.store").URL()
		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, data)
	}
}

func createTables()  {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
);`
	_, err := db.Exec(createArticlesSQL)
	checkError(err)
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main()  {
	initDB()
	createTables()
	
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", aritlcesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", aritlcesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles/create", aritlcesCreateHandler).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles", aritlcesStoreHandler).Methods("POST").Name("articles.store")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.Use(forceHTMLMiddleware)

	//homeURL, _ := router.Get("home").URL()
	//fmt.Println("homeURL: ", homeURL)
	//articleURL, _ := router.Get("articles.show").URL("id", "23")
	//fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
