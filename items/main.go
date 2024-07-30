package main

import (
	"log"

	"net/http"

	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	httpAddr = ":8080"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=item port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected to database")

	db.AutoMigrate(&Item{})

	// // Create
	// db.Create(&Item{Name: "D42", Price: 100})

	// // Read
	// var product Item
	// db.First(&product, 1) // find product with integer primary key
	// // db.First(&product, "code = ?", "D42") // find product with code D42

	// log.Println(product)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	handler := NewHandler(db)
	handler.Register(r)
	srv := &http.Server{
		Handler: r,
		Addr:    httpAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server is running at", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Products"))
}
