package main

import "net/http"

// Server 是http Server的顶级抽象
type Server interface {
	//Route 设置一个路由
	Route(pattern string, handlerFunc http.HandlerFunc)

	//Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer represents the HTTP server
type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, HandlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, HandlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewSdkHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
