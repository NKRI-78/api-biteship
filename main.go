package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"superapps/controllers"
	helper "superapps/helpers"
	middleware "superapps/middlewares"
	"superapps/services"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	services.InitDBs()

	if err != nil {
		helper.Logger("error", "Error getting env")
	}

	router := mux.NewRouter()

	router.Use(middleware.CorsMiddleware)

	// Check if the directory exists, create if it doesn't
	errMkidr := os.MkdirAll("public", os.ModePerm) // os.ModePerm ensures directory is created with the correct permissions
	if errMkidr != nil {
		log.Fatalf("Failed to create or access directory: %v", err)
	}

	// Open the public directory
	dir, err := os.Open("public")
	if err != nil {
		log.Fatalf("Failed to open public directory: %v", err)
	}
	defer dir.Close()

	// Read the directory contents
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalf("Failed to read directory contents: %v", err)
	}

	// Loop through each file in the directory
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			// Define static and public paths
			staticPath := "/" + fileInfo.Name() + "/"
			publicPath := "./public/" + fileInfo.Name() + "/"

			log.Printf("Serving static files from %s at %s", publicPath, staticPath)

			// Register (override if already exists) the route to serve static content
			router.PathPrefix(staticPath).Handler(http.StripPrefix(staticPath, http.FileServer(http.Dir(publicPath))))
		}
	}

	// Courier
	router.HandleFunc("/api/v1/courier-list", controllers.CourierList).Methods("GET")
	router.HandleFunc("/api/v1/courier-rate-list", controllers.CourierRateList).Methods("GET")
	router.HandleFunc("/api/v1/create-location", controllers.CreateLocation).Methods("POST")
	router.HandleFunc("/api/v1/rate-by-coordinate", controllers.RateByCoordinate).Methods("POST")

	// Tracking
	router.HandleFunc("/api/v1/tracking", controllers.Tracking).Methods("POST")

	// Order
	router.HandleFunc("/api/v1/order-by-coordinate", controllers.OrderByCoordinate).Methods("POST")
	router.HandleFunc("/api/v1/order-info", controllers.OrderInfo).Methods("POST")

	// PPOB
	router.HandleFunc("/api/v1/ppob/transaction-list", controllers.TransactionListPPOB).Methods("GET")

	// Payment
	router.HandleFunc("/api/v1/payment/transaction-list", controllers.TransactionListPayment).Methods("GET")

	portEnv := os.Getenv("PORT")
	port := ":" + portEnv

	fmt.Println("Starting server at", port)

	server := &http.Server{
		Addr:              port,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	errListenAndServe := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", errListenAndServe)
	}
}
