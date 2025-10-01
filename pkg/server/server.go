package server

import (
	"go_final_project/pkg/api"
	"net/http"
	"os"
)

func Run() {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	http.Handle("/", http.FileServer(http.Dir("./web")))
	api.Init()
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
