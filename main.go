package main

import "restful/router"

func main(){
	router := router.InitRouter()

	router.Run(":9000")
}