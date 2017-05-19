package main

// GetRouter creates a router and registers all the routes for the
import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// service and returns it.
func GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = PanicHandler

	// All GET routes
	router.GET("/hello", GreetUser)

	return router
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	log.Printf("Recovering from panic, Reason: %+v", c.(error))
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
