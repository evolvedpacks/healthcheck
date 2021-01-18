package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/evolvedpacks/healthcheck/internal/pinger"
)

var (
	flagAddr             = flag.String("addr", "", "Address and port of the Minecraft server")
	flagSilent           = flag.Bool("silent", false, "Whether to print console messages")
	flagValidateResponse = flag.Bool("validateResponse", true, "Wether to validate if the repsonse package is empty")
)

func main() {
	flag.Parse()

	if *flagAddr == "" {
		exit(2, "Error: Address must be provided. Use the -h flag for help.")
	}

	pong, err := pinger.PingOnce(*flagAddr)
	if err == pinger.ErrEmptyPong && !*flagValidateResponse {
		exitf(0, `{"result": "empty response"}`)
	}
	if err != nil {
		exitf(1, "Error: %s", err.Error())
	}

	data, _ := json.MarshalIndent(pong, "", "  ")
	exit(0, string(data))
}

func exit(code int, msg string) {
	if !*flagSilent {
		fmt.Println(msg)
	}
	os.Exit(code)
}

func exitf(code int, msg string, args ...interface{}) {
	exit(code, fmt.Sprintf(msg, args...))
}
