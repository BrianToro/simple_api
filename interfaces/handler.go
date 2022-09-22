package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BrianToro/simple_api/application"
	"github.com/BrianToro/simple_api/domain/models"
	"github.com/BrianToro/simple_api/infrastructure/migrations"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Run(port int) error {
	log.Printf("Server running at port: %d", port)
	server := http.Server{Addr: fmt.Sprintf(":%d", port), Handler: Routes()}
	return server.ListenAndServe()
}

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet)
	r.HandleFunc("/users", createUser).Methods(http.MethodPost)
	r.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", getUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", deleteUser).Methods(http.MethodDelete)

	// Migrations
	r.HandleFunc("/migrations", createMigrations).Methods(http.MethodPost)
	return r
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode("{\n  \"message\": \"Simple Api is running\"\n}")
	if err != nil {
		log.Panicln(err)
	}
}

func createMigrations(_ http.ResponseWriter, _ *http.Request) {
	err := migrations.Migrate()
	if err != nil {
		log.Panicln(err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	user, err := application.GetUser(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Panicln(err)
		}
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Panicln(err)
	}
}

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := application.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Panicln(err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser, err := application.Create(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		log.Panicln(err)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	err := application.Delete(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Panicln(err)
		}
	}
	w.WriteHeader(http.StatusOK)
	return
}
