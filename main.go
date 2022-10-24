// untuk menambah module lain gunakan "go get <link-module>"
// untuk upgrade module, tinggal ubah versi di go.mod lalu jalankan jalankan "go get"

package main

import (
	"fmt"

	gohello "github.com/orangdong/go-hello"
)

func main() {
	fmt.Println(gohello.SayHello())
}
