package handlers

import (
	"cars/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

func CarHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.Trim(entity, "/")

	switch r.Method {
	case "POST":
		if entity == "" {
			createCar(w, r)
		} else {
			http.Error(w, "Incorrect post request", http.StatusBadRequest)
		}

	case "GET":
		if entity == "" {
			http.Error(w, "We dont support this api", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			getCar(w, r, id)

		}
	case "DELETE":
		if entity == "" {
			http.Error(w, "We dont support this api", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			deleteCar(w, r, id)

		}

	case "PUT":
		if entity == "" {
			http.Error(w, "We dont support this api", http.StatusBadRequest)
		} else {
			id, _ := strconv.Atoi(entity)
			updateCar(w, r, id)
		}

	}

}

// {
// 	"name": "xw1"
// 	"model"
// }

func createCar(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "incorrect json input", http.StatusBadRequest)
		return
	}

	car.Insert()

	fmt.Println("Car saved to the inventory with thew id: ", car.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func updateCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{ID: id}
	if err := car.Get(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("car not found with the id: ", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "incorrect json input", http.StatusBadRequest)
		return
	}

	car.Update()

	fmt.Println("car updated with the id: ", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func getCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{ID: id}
	if err := car.Get(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("car not found with the id: ",id)
		return
	}

	fmt.Println("car found with the id:", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func deleteCar(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{ID: id}
	
    car.Delete()

	fmt.Println("car deleted with the id: ", id)



	w.WriteHeader(http.StatusNoContent)

}
