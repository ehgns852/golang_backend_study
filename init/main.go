package main

import (
	"flag"
	"fmt"
	"golang_backend_study/init/cmd"
	"os"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	fmt.Println(os.Getwd())
	cmd.NewCmd(*configPathFlag)
}
