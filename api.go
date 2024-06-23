package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	address  string
	database Database
}
type APIError struct {
	Error string
}

// For the JSON response
func ResponseJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	response := json.NewEncoder(w).Encode(data)
	return response
}

// Server handlers - Endpoints
type signature func(http.ResponseWriter, *http.Request) error

func httpHandleConverter(f signature) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			ResponseJSON(w, http.StatusBadRequest, APIError{
				err.Error(),
			})
		}
	}
}

// NewAPIServer creates a new instance of the API server.
//
// listen: string representing the address to listen on.
// Returns a pointer to the newly created Server instance.
func NewAPIServer(listen string, dbContext Database) *Server {
	// Create a new Server instance and set its address to the provided listen address.
	return &Server{
		address:  listen,
		database: dbContext,
	}
}

func (s *Server) Serve() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/products", httpHandleConverter(s.rootHandler))

	log.Println("Serve in", s.address)

	return http.ListenAndServe(s.address, router)
}

// Server handlers - Endpoints
func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.listProductsHandler(w, r)
	}
	if r.Method == http.MethodPost {
		return s.createProductHandler(w, r)
	}
	if r.Method == http.MethodPut {
		return s.updateProductHandler(w, r)
	}
	if r.Method == http.MethodDelete {
		return s.deleteProductHandler(w, r)
	}

	return nil
}

func (s *Server) createProductHandler(w http.ResponseWriter, r *http.Request) error {
	createProduct := new(NewProductRequest)
	if err := json.NewDecoder(r.Body).Decode(createProduct); err != nil {
		return err
	}

	// Save the product to the database
	product := NewProduct(createProduct.Name,
		createProduct.Description,
		createProduct.Price,
		createProduct.Quantity,
		createProduct.DeliveryDays,
	)
	if err := s.database.CreateProduct(product); err != nil {
		return err
	}

	return ResponseJSON(w, http.StatusCreated, product)
}

func (s *Server) updateProductHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) deleteProductHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) getProductHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	return ResponseJSON(w, http.StatusOK, vars)
}

func (s *Server) listProductsHandler(w http.ResponseWriter, r *http.Request) error {
	products, err := s.database.ListProducts()

	if err != nil {
		return err
	}
	return ResponseJSON(w, http.StatusOK, products)
}
