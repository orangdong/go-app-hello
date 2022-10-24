// untuk menambah module lain gunakan "go get <link-module>"

package main

import (
	"fmt"

	gohello "github.com/orangdong/go-hello"
)

func main() {
	fmt.Println(gohello.SayHello())
}
