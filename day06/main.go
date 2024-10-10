package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// 定义文件路径
	filePath := "C:\\Users\\dell\\log.txt"

	// 创建http服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			http.NotFound(w, r)
			return
		}
		if r.Method == http.MethodGet {
			// 获取当前时间
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			// 构建要写入的字符串
			content := fmt.Sprintf("pressuretest%s\n", currentTime)

			// 打开文件，准备追加内容
			file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// 写入文件

			_, err = io.WriteString(file, content)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// 响应请求
			w.Write([]byte("Data written to file."))
		} else {
			// 如果不是GET请求，返回405 Method Not Allowed
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 启动服务端
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
