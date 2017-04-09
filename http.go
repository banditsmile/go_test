package main

import (
	"log"
	"net/http"
	"./routers"
)

func main(){
	routers.Index()
	err:=http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServer:", err)
	}
}