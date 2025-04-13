package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"superapps/controllers"
	helper "superapps/helpers"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helper.Logger("error", "Error getting env")
	}

	router := mux.NewRouter()

	// router.Use(middleware.JwtAuthentication)

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

	// Inisialisasi rate limiter: 2 permintaan per menit
	// rateLimiter := middleware.NewRateLimiter(2, 1)

	// // Auth
	// router.Handle("/api/v1/login", rateLimiter.LimitMiddleware(http.HandlerFunc(controllers.Login))).Methods("POST")

	// Courier
	router.HandleFunc("/api/v1/courier-list", controllers.CourierList).Methods("GET")
	router.HandleFunc("/api/v1/create-location", controllers.CreateLocation).Methods("POST")

	// Courier

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

	// errs := http.ListenAndServe(port, router)
	// if errs != nil {
	// 	fmt.Println("Error starting server:", errs)
	// }
}
