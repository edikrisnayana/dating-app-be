package main

import (
	"datingAppBE/route"
)

func main() {
	router := route.GetRouter()

	router.Start("localhost:8080")
}
