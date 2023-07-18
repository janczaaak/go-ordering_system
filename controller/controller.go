package controller

import (
	"encoding/json"
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
		if err != nil{
			log.Fatal(err)
		}
		orders = append(orders,order)

		response.Status = 200
		response.Message = "Succes"
		response.Data = orders
		
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Acces-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}

}

func AddOrder()    {}
func UpdateOrder() {}
func DeleteOrder() {}
