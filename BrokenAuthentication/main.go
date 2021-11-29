package main

import "BrokenAuthentication/router"

func main(){
	router.InitializeRouter().Run(":4444")
}