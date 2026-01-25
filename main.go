package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int     `json:"id"`
	Nama  string  `json:"nama"`
	Harga float64 `json:"harga"`
	Stok  int     `json:"stok"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var produk = []Produk{
	{ID: 1, Nama: "Laptop", Harga: 5000000, Stok: 10},
	{ID: 2, Nama: "Smartphone", Harga: 2000000, Stok: 25},
	{ID: 3, Nama: "Tablet", Harga: 3000000, Stok: 15},
	{ID: 4, Nama: "Headset", Harga: 25000, Stok: 30},
}

var categories = []Category{
	{ID: 1, Name: "Elektronik", Description: "Produk elektronik"},
	{ID: 2, Name: "Peralatan", Description: "Produk peralatan"},
	{ID: 3, Name: "Aksesori", Description: "Produk aksesori"},
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Produk not found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "Produk not found", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Produk deleted"})
			return
		}
	}
	http.Error(w, "Produk not found", http.StatusNotFound)

}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, p := range categories {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id
			categories[i] = updateCategory
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategory)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, p := range categories {
		if p.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted"})
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)

}

func main() {
	// GET Produk localhost:8080/api/produk/{id}
	// PUT Produk localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})

	// GET Produk localhost:8080/api/produk
	// POST Produk localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var newProduk Produk
			err := json.NewDecoder(r.Body).Decode(&newProduk)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newProduk.ID = len(produk) + 1
			produk = append(produk, newProduk)
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newProduk)
		}
	})

	// GET Category localhost:8080/api/categories/{id}
	// PUT Category localhost:8080/api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryByID(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	// GET Category localhost:8080/api/categories
	// POST Category localhost:8080/api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		} else if r.Method == "POST" {
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newCategory)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy", "message": "API Running!"})
	})

	fmt.Println("▶️  Starting server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
