//curl -i -X POST -F file=@maxresdefault.jpg   localhost:8080/upload
package main

import (
    "fmt"
	"log"
	"net/http"
	"time"
	"github.com/gomodule/redigo/redis"
    "github.com/gorilla/mux"
    "os"
    "io"
)

var pool *redis.Pool
func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Status OK"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
    file, handler, err := r.FormFile("file")
    fmt.Println(file)
    fileName := r.FormValue("file")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    _, _ = io.WriteString(w, "File "+fileName+" Uploaded successfully")
    _, _ = io.Copy(f, file)
}


func main() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
    r := mux.NewRouter()
    r.HandleFunc("/status", get).Methods(http.MethodGet)
    r.HandleFunc("/upload", UploadFile).Methods(http.MethodPost)
    r.HandleFunc("/", notFound)
    log.Fatal(http.ListenAndServe(":8080", r))
}
