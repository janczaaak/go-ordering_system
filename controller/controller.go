package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/janczaaak/go-ordering_system/config"
	"github.com/janczaaak/go-ordering_system/model"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	var orders []model.Order
	var response model.Response

	db := config.ConnectDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err := rows.Scan(&order.ID, &order.Product, &order.Amount, &order.Status, &order.Client.FirstName, &order.Client.LastName, &order.Client.PhoneNumb, &order.Client.Email)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}
	response.Status = 200
	response.Message = "Succes"
	response.Data = orders

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Fatal(err)
	}
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.ConnectDb()
	defer db.Close()

	Product := r.FormValue("Product")
	Amount := r.FormValue("Amount")
	Status := r.FormValue("Status")
	FirstName := r.FormValue("FirstName")
	LastName := r.FormValue("LastName")
	PhoneNumb := r.FormValue("PhoneNumb")
	Email := r.FormValue("Email")

	query := "INSERT INTO orders(Id, Product, Amount, Status, Firstname, Lastname, Phonenumb, Email) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, 0, Product, Amount, Status, FirstName, LastName, PhoneNumb, Email)
	if err != nil {
		log.Fatal(err)
	}

	response.Status = 200
	response.Message = "Succes"
	fmt.Print("Instert data")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Fatal(err)
	}
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {

	var response model.Response

	db := config.ConnectDb()
	defer db.Close()

	ID := r.FormValue("ID")
	Product := r.FormValue("Product")
	Amount := r.FormValue("Amount")
	Status := r.FormValue("Status")
	FirstName := r.FormValue("FirstName")
	LastName := r.FormValue("LastName")
	PhoneNumb := r.FormValue("PhoneNumb")
	Email := r.FormValue("Email")

	updateOrd, err := db.Prepare("UPDATE orders SET Product=?, Amount=?, Status=?, FirstName=?, LastName=?, PhoneNumb=?, Email=? WHERE ID=?")
	if err != nil {
		panic(err)
	}
	_, err = updateOrd.Exec(Product, Amount, Status, FirstName, LastName, PhoneNumb, Email, ID)
	if err != nil{
		panic(err)
	}

	response.Status = 200
	response.Message = "Updated"
	fmt.Print("Update data")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Fatal(err)
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {

	var response model.Response

	db := config.ConnectDb()
	defer db.Close()
	ID := r.FormValue("ID")
	_, err := db.Exec("DELETE FROM orders WHERE id=?", ID)
	if err != nil {
		panic(err)
	}

	response.Status = 200
	response.Message = "Succes"
	fmt.Print("Delete data")
	err = json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Fatal(err)
	}
}
