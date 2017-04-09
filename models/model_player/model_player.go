package model_player

import (
	"log"
	"../../../github.com/seefan/gossdb"
)

var ssdbClient *gossdb.Client


func init(){
	pool, ssdbErr := gossdb.NewPool(&gossdb.Config{
		Host:             "127.0.0.1",
		Port:             6380,
		MinPoolSize:      5,
		MaxPoolSize:      50,
		AcquireIncrement: 5,
	})
	if ssdbErr != nil {
		log.Fatal(ssdbErr)
		return
	}
	ssdbClient, err := pool.NewClient()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer ssdbClient.Close()
}
func Player_info(hashName,key string)(gossdb.Value,error){
	//设置10 秒过期
	ssdbClient.Set(hashName,key,10)
	return ssdbClient.Get(hashName)
}