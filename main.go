package main

import "net/http"

func main() {
	r := SetRouter()

	http.ListenAndServe(":3456", r)
}
