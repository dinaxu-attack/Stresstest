package main

import (
	"fmt"
	"os"

	"github.com/dinaxu-attack/Stresstest/app"
)

func help() {
	fmt.Println("Stresstest <URL> <COUNT>\nExample:\nStresstest https://tls.mrrage.xyz 1000")
}

func main() {

	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}

	app.MainStress(os.Args[1], os.Args[2])
}
