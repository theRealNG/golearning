package main

import (
	"fmt"

	"github.com/theRealNG/exampleHTTPRequests/reqres"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	user := reqres.User{
		Name: "Harry Potter",
		Job:  "Auror",
	}

	reqres.Create(&user)

	fmt.Println(user)
}
