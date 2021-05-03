package main

import (
	"net/http"

	"github.com/lunarxlark/learn-go-with-tests/dependency_injection"
)

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(dependency_injection.MyGreeterHandler))
}