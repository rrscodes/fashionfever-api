package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rrscodes/fashionfever-api/pkg/models"
	"github.com/rrscodes/fashionfever-api/pkg/utils"
)

var NewDress models.Dress

func GetAllDresses(w http.ResponseWriter, r *http.Request) {
	newDress := models.GetAllDresses()
	res, _ := json.Marshal(newDress)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDressById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	dressDetails, _ := models.GetDressById(ID)
	res, _ := json.Marshal(dressDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateDress(w http.ResponseWriter, r *http.Request) {
	CreateDress := &models.Dress{}
	utils.ParseBody(r, CreateDress)
	d := CreateDress.CreateDress()
	res, _ := json.Marshal(d)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteDress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	d := models.DeleteDress(ID)
	res, _ := json.Marshal(d)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateDress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	UpdateDress := &models.Dress{}
	utils.ParseBody(r, UpdateDress)
	dressDetails, db := models.GetDressById(ID)
	if UpdateDress.Label != "" {
		dressDetails.Label = UpdateDress.Label
	}
	if UpdateDress.Brand != "" {
		dressDetails.Brand = UpdateDress.Brand
	}
	if UpdateDress.Size != "" {
		dressDetails.Size = UpdateDress.Size
	}
	if UpdateDress.Color != "" {
		dressDetails.Color = UpdateDress.Color
	}
	db.Save(&dressDetails)
	res, _ := json.Marshal(dressDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
