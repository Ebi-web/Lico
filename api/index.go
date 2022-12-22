package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. /helloというエンドポイントを登録する
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 2. クライアントからのリクエストを処理する
		fmt.Fprintf(w, "Hello, World!")
	})

	// 3. サーバーを起動する
	http.ListenAndServe(":8080", nil)
}
