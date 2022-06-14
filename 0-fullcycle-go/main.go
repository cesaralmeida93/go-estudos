package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.ListenAndServe(":8887", nil)
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint
}