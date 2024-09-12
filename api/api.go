package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Bir HTTP yanıtını JSON formatında gönderir.Bir http yanıtına JSON veri ekler ve uygun başlıkları ayarlar.
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	//Header'a content-type ekler, Yanıtın JSON formatında olduğunu belirtir ve istemcinin bu yanıtı JSOn olarak işlemesini sağlar
	w.Header().Set("Content-Type", "application/json")
	//NewEncoder: Bir JSON encoder(kodlayıcı) oluşturur.Bu encoder, w üzerine JSON verisini yazacaktır
	//Encode:v değerini JSON formatına dönüştürür ve bu JSON verisini HTTP yanıtına yazar
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

// HandlerFunc: Http işlevlerinin işlenmesini sağlayan bir türdür.
func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	//Gelen HTTP isteklerini işleyen anonim bir işlevdir.
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		//Fonksiyondaki listenAddr'yi listenAddr ile doldurur.
		listenAddr: listenAddr,
	}
}
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))
	log.Println("JSON API server running on board", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// API server üzerindeki bir nesne üzerinde çalışacağını gösterir.
// handleAccount'tan sonrakiler parametredir.
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// http.ResponseWriter : Http yanıtı oluşturmak için kullanılır.
// r. *http.Request: Gelen HTTP isteğinin detaylarını içerir
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil

}
