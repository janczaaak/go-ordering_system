package model

type Order struct {
	ID      string `form:"id" json:"id"`
	Product string `form:"product" json:"product"`
	Amount  int `form:"amount" json:"amount"`
	Status  string `form:"status" json:"status"`
	Client  struct {
		FirstName string `form:"firstname" json:"firstname"`
		LastName  string `form:"lastname" json:"lastname"`
		PhoneNumb string `form:"phonenumb" json:"phonenumb"`
		Email     string `form:"email" json:"email"`
	}
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Order
}