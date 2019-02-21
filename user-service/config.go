package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct{}

// AppConfig setup to get access database configurations using envVar
type AppConfig struct {
	DB []dbconf `json:"DB"`
}

type dbconf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Name     string `json:"DBName"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

// SetEnvVars to set all configs to the env vars
func (conf *config) SetEnvVars() error {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := AppConfig{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	dbc := configuration.DB[0]
	os.Setenv("DB_HOST", dbc.Host)
	os.Setenv("DB_USER", dbc.User)
	os.Setenv("DB_NAME", dbc.Name)
	os.Setenv("DB_PASSWORD", dbc.Password)
	os.Setenv("DB_PORT", dbc.Port)

	return nil
}
