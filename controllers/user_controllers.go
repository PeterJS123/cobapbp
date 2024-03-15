package controllers

import (
	"encoding/json"
	"fmt"
	m "latihan/models"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	// m.CobaHitung()
	query := "SELECT * FROM users"
	name := r.URL.Query()["name"]
	age := r.URL.Query()["age"]
	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name='" + name[0] + "'"
	}
	if age != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age='" + age[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType); err != nil {
			log.Println(err)
			return
		} else {
			users = append(users, user)
		}
	}
	w.Header().Set("content-Type", "application/json")
	var response m.UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	// m.CobaHitung()
	query := "SELECT * FROM users"
	id := r.URL.Query()["id"]
	name := r.URL.Query()["name"]
	if id != nil {
		fmt.Println(id[0])
		query += " WHERE id='" + id[0] + "'"
	}
	if name != nil {
		if id[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " name='" + name[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var products m.products
	var productss []m.products
	for rows.Next() {
		if err := rows.Scan(&products.ID, &products.Name, &products.Age, &products.Address, &products.productsType); err != nil {
			log.Println(err)
			return
		} else {
			users = append(users, products)
		}
	}
	w.Header().Set("content-Type", "application/json")
	var response m.UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)
}
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	// m.CobaHitung()
	query := "SELECT * FROM users"
	id := r.URL.Query()["id"]
	userid := r.URL.Query()["userid"]
	if id != nil {
		fmt.Println(id[0])
		query += " WHERE id='" + id[0] + "'"
	}
	if userid != nil {
		if id[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " userid='" + userid[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var transaction m.transactions
	var transactions []m.transactions
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.Name, &transaction.Age, &transaction.Address, &transaction.transactionType); err != nil {
			log.Println(err)
			return
		} else {
			users = append(users, transaction)
		}
	}
	w.Header().Set("content-Type", "application/json")
	var response m.TransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = transactions
	json.NewEncoder(w).Encode(response)
}
