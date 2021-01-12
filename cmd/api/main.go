package main

import "ctag-go-learning/router"

func main() {
	err := router.GetRouter().Run(":3000")
	if err != nil {
		panic(err)
	}
}