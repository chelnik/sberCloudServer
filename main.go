package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/chelnik/sbercloudServer/config"
	"github.com/gorilla/mux"
)

const pathConfig string = "/Users/vadimcelnik/go/sberCloudServer/config.yaml"

func (d *Data) middlewareHandler(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middlewareHandler", r.URL.Path)
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("[%s] %s %s %s\n\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
		d.writeInFileMap()
	})
}

// writeInFileMap записывает данные в файл что-бы видеть что находится в мапе
func (d *Data) writeInFileMap() {
	file, err := os.Create("text.txt")
	if err != nil {
		log.Println("file open error")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("error when closing the file")
		}
	}(file)
	for i, item := range d.Mapa {
		_, err = file.Write([]byte(fmt.Sprintf("%s%v\n", i, item)))
		if err != nil {
			log.Println("file write error")
		}
	}
}
func main() {
	address := config.NewPointerAddress() // работа с конфигом
	address.LoadConfig(pathConfig)
	mapa := &Data{make(map[string]StructForJson)}

	router := mux.NewRouter() // горилла

	mapa.addHandlers(router)
	coveredMux := mapa.middlewareHandler(router)
	fmt.Printf("server listen at http://localhost%s\n", address.Port)
	err := http.ListenAndServe(address.Port, coveredMux)
	if err != nil {
		log.Println("error with server")
		debug.PrintStack() // выводит стек трейс
	}
}
