package main

import (
	"BrokenAuthentication/router"
)

func main(){

	rout := router.InitializeRouter()
	rout.Run(":4444")
}