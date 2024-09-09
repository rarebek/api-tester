package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"fuzzer/cmd"
	"fuzzer/config"
)

func main() {
	cmd.Execute()

	var cfg config.Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal(err)
	}

}
