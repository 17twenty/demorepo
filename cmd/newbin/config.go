package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type databaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

type generalConfig struct {
	Port int `json:"port"`
}

// config is the main configuration for the project.
// we initialise local dev settings here which we override at
// runtime for prod/staging
var config = struct {
	Database databaseConfig `json:"database"`
	General  generalConfig  `json:"general"`
}{
	Database: databaseConfig{
		User:     "localuser",
		Password: "password",
		Host:     "",
		Port:     5432,
		Name:     "dockerdb",
	},
	General: generalConfig{
		Port: 8080,
	},
}

func writeDefaultConfig(location string) {
	f, err := os.Create(location)
	if err != nil {
		log.Fatalln("Couldn't open", location)
	}

	data, _ := json.MarshalIndent(config, "", "  ")
	_, err = f.Write(data)
	if err != nil {
		log.Fatalln("Couldn't write to", location)
	}
}

func loadConfig(location string) {
	raw, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatalln("Couldn't open", location)
	}

	err = json.Unmarshal(raw, &config)
	if err != nil {
		log.Fatalln("Couldn't understand config in", location, "-", err)
	}
}
