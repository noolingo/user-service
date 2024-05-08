package main

import (
	"fmt"
	"log"
	"os"

	"github.com/noolingo/user-service/internal/app"
)

// не пушить!!!!
const configPath = "configs/config.yml"

func main() {
	err := app.Run(configPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("service work!!!\n")
	os.Exit(0)
}
