package main

import (
	"fmt"
	"github.com/sivaramsajeev/go-starter/dirs"
	"github.com/sivaramsajeev/go-starter/internal/auth"
)

func main() {
	fmt.Println("Gopher z Just starting out :)")
	fmt.Println(auth.GetAuth())
	fmt.Println(dirs.AddonPrint())
}