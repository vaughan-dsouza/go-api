package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/vaughan-dsouza/go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.newRouter()
	handlers.Handler(r)

	fmt.Println("starting GO Api service...")
	fmt.Println(`
		   ______      ___    ____  ____
		  / ____/___  /   |  / __ \/  _/
		 / / __/ __ \/ /| | / /_/ // /  
		/ /_/ / /_/ / ___ |/ ____// /   
		\____/\____/_/  |_/_/   /___/   								
	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil{
		log.Error(err)
	}
}