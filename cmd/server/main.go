package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ravigill3969/Log-Aggregator-status-go/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Unable to load .env file " + err.Error())
		return
	}

	_, err := os.Stat("logs.txt")

	if os.IsNotExist(err) {
		file, err := os.Create("logs.txt")
		if err != nil {
			fmt.Println("Unable to create new file: " + err.Error())
			return
		}
		defer file.Close()
		fmt.Println("logs.txt created")
	} else {
		fmt.Println("logs.txt already exists")
	}

	mux := http.NewServeMux()

	routes.RoutesForAggregation(mux)

	PORT := ":" + os.Getenv("PORT")

	if PORT == "" {
		fmt.Println("PORT not provided")
	}

	fmt.Printf(
		"Server is running on PORT %s http://localhost%s\n",
		PORT, PORT,
	)

	if err := http.ListenAndServe(PORT, mux); err != nil {
		log.Fatal("Unable to start server")
	}
}
