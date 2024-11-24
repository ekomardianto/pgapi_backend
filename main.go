package main

import (
	"eko/api-pg-bpr/middlewares"
	"eko/api-pg-bpr/models"
	"eko/api-pg-bpr/routes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("selamat datang ekoo")
	// Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("load file env berhasil")
	// script mendapatkan port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Get the address the listener is listening on
	addr := listener.Addr().(*net.TCPAddr)

	// Print the port number
	fmt.Printf("Server is starting at port %d\n", addr.Port)
	// end script

	models.ConnectDatabase()

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	//Routing Auth
	routes.AuthRoutes(r)

	//routing CSRF
	routes.CSRFRoutes(api)

	//Routing User
	routes.UserRoutes(api)

	//Routing RoleUser
	routes.RoleUserRoutes(api)

	//Routing Instansi
	routes.InstansiRoutes(api)

	//Routing Perusahaan
	routes.PerusahaanRoutes(api)

	//Routing Ipaymu
	routes.Ipayroutes(api)

	//bungkus ke dalam middlewares semua prefix /api
	api.Use(middlewares.JWTMiddleware)
	// Konfigurasi middleware CORS
	var feURL = os.Getenv("FRONTEND_URL")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{feURL}, // Ubah sesuai dengan origin aplikasi frontend Anda
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Origin", "Accept", "Cookie", "cache-control", "expires", "X-CSRF-TOKEN"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}
