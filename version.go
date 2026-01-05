package main

import "fmt"

var (
	AppName    = "Biji"
	Version    = "v0.1.0"
	Author     = "dhy"
	Purpose    = "Defensive Web Risk & Traffic Monitoring Tool"
)

func PrintVersion() {
	fmt.Println(AppName)
	fmt.Println("Version :", Version)
	fmt.Println("Author  :", Author)
	fmt.Println("Purpose :", Purpose)
	fmt.Println("License :", "Educational / White-Hat Use")
}
