package main

import (
	"fmt"

	"github.com/devnjw/go-study/module/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}
