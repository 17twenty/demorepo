package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("Starting newbin")

	defaultLocation := flag.String("default", "", "location to write a default configuration to (this will overwrite an existing file at this location)")
	configLocation := flag.String("config", "", "JSON config file to load")

	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	if *defaultLocation == "" && *configLocation == "" {
		log.Println("Using default config:")
		data, _ := json.MarshalIndent(config, "", "  ")
		// No config file? Write to stdout to
		// highlight it shouldn't be compiled with secrets
		io.Copy(os.Stdout, bytes.NewReader(data))
		fmt.Printf("\n")
	} else if *defaultLocation != "" {
		writeDefaultConfig(*defaultLocation)
		os.Exit(0)
	} else if *configLocation != "" {
		loadConfig(*configLocation)
	}

	log.Println("Here I am")
}
