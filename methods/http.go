package main

import (
	"fmt"
	"net/http"
)

// 定义价格类型
type price float64

func (p price) String() string {
	return fmt.Sprintf("%.2f¥", p)
}

// 使用 map 模拟数据库
type database map[string]float64

// 定义 http.Handler 接口
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/list":
		for key, val := range db {
			fmt.Fprintf(w, "%s：%s\n", key, price(val))
		}
	case "/price":
		// 获取 URL 参数 item
		item := r.URL.Query().Get("item")
		// 在 db 中查询数据，如果不存在返回错误
		if value, ok := db[item]; ok {
			fmt.Fprintf(w, "%s\n", price(value))
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %s\n", item)
		}
	default:
		// header := w.Header()
		// header["Set-Cookies"] = []string{"token=werwewrwerwr23342;path=/;"}
		msg := fmt.Sprintf("no such page: %s\n", path)
		http.Error(w, msg, http.StatusNotFound)

	}
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for key, val := range db {
		fmt.Fprintf(w, "%s：%s\n", key, price(val))
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数 item
	item := r.URL.Query().Get("item")
	// 在 db 中查询数据，如果不存在返回错误
	if value, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price(value))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n", item)
	}
}
