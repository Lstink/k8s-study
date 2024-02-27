package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 响应显示的版本号
	var version = "v1.0.0"
	// 端口号
	var port = 8080

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		_, _ = w.Write([]byte("this is version " + version + ", hostname: " + hostname))
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("this is health check response"))
	})

	fmt.Println("Server is running at :8080")

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("server is error: ", err.Error())
	}
}
