package main

import (
	"os"
	"fmt"
	"alfred-jenkins/jenkins"
)

func main(){

	url := os.Args[1:3][0]

	h := jenkins.Handler{url}

	response := h.GetStatus("test")

	fmt.Println(response);

}
