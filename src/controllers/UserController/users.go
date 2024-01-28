package usercontroller

import (
	"blanja_api/src/helper"
	"blanja_api/src/middleware"
	models "blanja_api/src/models/UserModel"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SellerRegister(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:        input.Name,
			Email:       input.Email,
			Phonenumber: input.Phonenumber,
			Storename:   input.Storename,
			Password:    Password,
			Role:        "Seller",
		}
		models.PostUser(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Seller Account Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func CustomerRegister(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:        input.Name,
			Email:       input.Email,
			Phonenumber: "-",
			Storename:   "-",
			Password:    Password,
			Role:        "Customer",
		}
		models.PostUser(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Customer Account Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		ValidateEmail := models.FindEmail(&input)
		if len(ValidateEmail) == 0 {
			fmt.Fprintf(w, "Email not Found")
			return
		}
		var passwordSecond string
		for _, user := range ValidateEmail {
			passwordSecond = user.Password
		}
		if err := bcrypt.CompareHashAndPassword([]byte(passwordSecond), []byte(input.Password)); err != nil {
			fmt.Fprintf(w, "Password not Found")
			return
		}
		item := map[string]string{
			"Message": "HI," + input.Name + " as a " + input.Role,
		}
		var result, _ = json.Marshal(item)
		w.Write(result)
		return
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_users(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAllUser().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		return
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_user(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/user/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectUserById(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "DELETE" {
		models.DeleteUser(id)
		msg := map[string]string{
			"Message": "User Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Update_seller(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/update-seller/"):]

	if r.Method == "PUT" {
		var updateSeller models.User
		err := json.NewDecoder(r.Body).Decode(&updateSeller)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newSeller := models.User{
			Name:        updateSeller.Name,
			Email:       updateSeller.Email,
			Phonenumber: updateSeller.Phonenumber,
			Storename:   updateSeller.Storename,
			Password:    updateSeller.Password,
		}
		models.UpdateSeller(id, &newSeller)
		msg := map[string]string{
			"Message": "Seller Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Update_customer(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/update-customer/"):]

	if r.Method == "PUT" {
		var updateCustomer models.User
		err := json.NewDecoder(r.Body).Decode(&updateCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCustomer := models.User{
			Name:     updateCustomer.Name,
			Email:    updateCustomer.Email,
			Password: updateCustomer.Password,
		}
		models.UpdateCustomer(id, &newCustomer)
		msg := map[string]string{
			"Message": "Customer Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
