package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/blackviking27/go-data-management/pkg/models"
	"github.com/blackviking27/go-data-management/pkg/utils"
	"github.com/gorilla/mux"
)

var newData models.Data

func GetData(w http.ResponseWriter, r *http.Request) {
	newData := models.GetAllData()
	res, _ := json.Marshal(newData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDataById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataId := vars["dataId"]
	Id, err := strconv.ParseInt(dataId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", err)
	}
	dataDetails, _ := models.GetDataById(Id)
	res, _ := json.Marshal(dataDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	CreateData := &models.Data{}
	utils.ParseBody(r, CreateData)
	b := CreateData.CreateData()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataId := vars["dataId"]
	Id, err := strconv.ParseInt(dataId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing the data", err)
	}
	data := models.DeleteData(Id)
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	var updateData = &models.Data{}
	utils.ParseBody(r, updateData)
	vars := mux.Vars(r)
	dataId := vars["dataId"]
	Id, err := strconv.ParseInt(dataId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", err)
	}
	dataDetails, db := models.GetDataById(Id)
	if updateData.Name != "" {
		dataDetails.Name = updateData.Name
	}
	if updateData.Author != "" {
		dataDetails.Author = updateData.Author
	}
	if updateData.Publication != "" {
		dataDetails.Publication = updateData.Publication
	}
	db.Save(&dataDetails)
	res, _ := json.Marshal(dataDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
