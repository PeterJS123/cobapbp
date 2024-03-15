package models

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
type Transactions struct {
	ID        int    `json:"id"`
	UserID    string `json:"userid"`
	ProductID int    `json:"prodictid"`
	Quantity  string `json:"quantity"`
}
type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}
type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
