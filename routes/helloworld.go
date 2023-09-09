package routes

import (
	"github.com/gorilla/mux"
	"hello/handlers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	// 你可以在这里添加更多的路由
	return r
}
