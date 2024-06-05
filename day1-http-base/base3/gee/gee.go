package gee

import(
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter,*http.Request)

type Engine struct{
	router map[string]HandleFunc
}

func New() *Engine{
	return &Engine{router:make(map[string]HandleFunc)}
}

