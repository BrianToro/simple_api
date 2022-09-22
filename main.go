package main

import (
	"github.com/BrianToro/simple_api/interfaces"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}

	err = interfaces.Run(8085)
	if err != nil {
		log.Panic(err)
	}
}
