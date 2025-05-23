package main

import (
"fmt"
"log"
"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "아름다운가게 지원을 환영합니다!")
}

func main() {
http.HandleFunc("/", helloHandler)
log.Println("Go API 서버가 정상적으로 실행되었습니다! 포트: 8080")
log.Fatal(http.ListenAndServe(":8080", nil))
}
