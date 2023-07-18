package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/janczaaak/go-ordering_system/controller"
)


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/getOrders", controller.GetOrders).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234",  router))
}





