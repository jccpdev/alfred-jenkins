package main

import (
	"os"
	"fmt"
	"alfred-jenkins/jenkins"
)

func main(){

	url := os.Args[1:3][0]
	query := os.Args[1:3][1]

	h := jenkins.Handler{url}

	response := h.GetStatus(query)

	fmt.Println(response);

}
