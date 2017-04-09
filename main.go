package main

import (
	"log"
	"./models/model_player"
)


func main(){
	res,e :=model_player.Player_info("aa","bb")
	log.Println(res);
	log.Println(e);
}