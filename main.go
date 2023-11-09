package main

import (
	"fmt"
	"log"
	"mogoIn/router"
	"net/http"
)

func main(){
	fmt.Println("Welcome to netflix api")
	r := router.Mains() 

	log.Fatal(http.ListenAndServe(":8000", r))
	

}

