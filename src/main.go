//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	srv := NewSdkHttpServer("my-http-server")

	//注册路由
	srv.Route("/body/one", readBodyOnce)
	srv.Route("/url/query", queryParams)
	srv.Route("/wholeUrl", wholeUrl)
	srv.Route("/header", header)

	srv.Start("8080")
}
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	fmt.Fprintf(w, "read body data: %s \n", string(body))

	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	fmt.Fprintf(w, "read body data one more time :[%s] and read data length %d \n", string(body), len(body))
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Fprintf(w, "query params: %v \n", query)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, "%s", string(data))
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header: %v \n", r.Header)
}
