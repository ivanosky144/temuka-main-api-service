package main

import (
	"log"
	"net/http"

	router "github.com/temuka-api-service/api"
	"github.com/temuka-api-service/config"
	"gorm.io/gorm"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.OpenConnection()
	var db *gorm.DB = config.GetDBInstance()

	if config.Database == nil {
		log.Fatal("Database connection is nil")
	}

	config.InitRedis()
	config.InitS3()

	router := router.Routes(db)
	protectedRoutes := EnableCors(router)

	http.HandleFunc("/chat", config.HandleWebSocket)
	go config.RecentHub.Run()

	http.Handle("/", protectedRoutes)
	log.Println("Server is listening on port 3200")
	log.Fatal(http.ListenAndServe("localhost:3200", nil))
}