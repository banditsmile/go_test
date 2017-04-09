package routers

import (
	"../controllers/player"
	"net/http"
)
func Index(){
	http.HandleFunc("/player/",player.Hash_handler)
}

