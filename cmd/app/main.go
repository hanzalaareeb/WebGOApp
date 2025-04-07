package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper to read configuration
	initConfig()
	appName := viper.GetString("app_name")
	message := viper.GetString("welcome_message")
	fmt.Printf("[%s]: %s\n", appName, message)

	// Set up a simple HTTP server
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bye", byeHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bye World!")
}

func initConfig() {
	viper.SetConfigName("config") // name of config file
	viper.SetConfigType("yaml")   // REQUIRED
	viper.AddConfigPath(".")      // optionally loook for config in working dir

	err := viper.ReadInConfig() // find and read the config file
	if err != nil {
		log.Fatal("Error reading config file", err)
	}

}
