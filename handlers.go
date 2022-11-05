package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StructForJson структура для хранения запроса
type StructForJson struct {
	Service string `json:"service"`
	Data    []struct {
		Key1 string `json:"key1,omitempty"`
		Key2 string `json:"key2,omitempty"`
	} `json:"data"`
}

// Data для хранения config
type Data struct {
	Mapa map[string]StructForJson
}

func (d *Data) addHandlers(router *mux.Router) {
	router.HandleFunc("/config", d.createConfig).Methods("POST")
	router.HandleFunc("/delete", d.deleteConfig).Methods("POST")
	router.HandleFunc("/update", d.updateConfig).Methods("POST")

	router.HandleFunc("/config", d.getConfig).Methods("GET")
}

// createConfig создание config
func (d *Data) createConfig(w http.ResponseWriter, r *http.Request) {
	var item = new(StructForJson)
	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("error with readAll r.Body")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, item)
	if checkItem := d.Mapa[item.Service]; checkItem.Service != "" { // проверка на дублирование сервиса
		http.Error(w, "service already exists", http.StatusBadRequest)
		return
	}
	d.Mapa[item.Service] = *item // кладем указатель на новый элемент в мапу
	if err != nil {
		log.Println("error with Unmarshal")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("created"))
	if err != nil {
		return
	}
}

// getConfig получение config
func (d *Data) getConfig(w http.ResponseWriter, r *http.Request) {
	service := r.URL.Query().Get("service") // считывание параметра
	response := d.Mapa[service]
	if response.Service == "" {
		http.Error(w, "the data does not exist", http.StatusBadRequest)
		return
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (d *Data) deleteConfig(w http.ResponseWriter, r *http.Request) {
	service := r.URL.Query().Get("service") // считывание параметра
	response := d.Mapa[service]
	if response.Service == "" {
		http.Error(w, "the data does not exist", http.StatusBadRequest)
		return
	}
	delete(d.Mapa, service)
}

func (d *Data) updateConfig(w http.ResponseWriter, r *http.Request) {
	var item = new(StructForJson)
	service := r.URL.Query().Get("service") // считывание параметра
	response := d.Mapa[service]
	if response.Service == "" {
		http.Error(w, "the data does not exist", http.StatusBadRequest)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error with readAll r.Body")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, item)
	if err != nil {
		http.Error(w, "service already exists", http.StatusBadRequest)
		return
	}
	d.Mapa[service].Data[0], d.Mapa[service].Data[1] = item.Data[0], item.Data[1]
}
