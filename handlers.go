package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
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
	conn *pgx.Conn
}

func (d *Data) middlewareHandler(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("[%s] %s %s %s %s\n\n", r.Method, r.UserAgent(), r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func (d *Data) addHandlers(router *mux.Router) {
	router.HandleFunc("/config", d.createConfig).Methods("POST")
	router.HandleFunc("/config", d.deleteConfig).Methods("DELETE")
	router.HandleFunc("/config", d.updateConfig).Methods("PATCH")

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
	d.insertConfigDb(item)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("created"))
	if err != nil {
		return
	}
}

// getConfig получение config
func (d *Data) getConfig(w http.ResponseWriter, r *http.Request) {
	service := r.URL.Query().Get("service") // считывание параметра
	response := d.selectConfig(service)
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

	d.deleteConfigDb(service)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("deleted"))
	if err != nil {
		return
	}

}

func (d *Data) updateConfig(w http.ResponseWriter, r *http.Request) {
	var item = new(StructForJson)
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error with readAll r.Body")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(bytes, item)
	d.insertConfigDb(item)
}
