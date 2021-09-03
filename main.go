package main

import (
	"fmt"

	go_say_hello "github.com/abdul-local/go-say-hello/v2"
)

func main(){

	hello:=go_say_hello.SayHello("abdul")

	fmt.Println(hello)
}