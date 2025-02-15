package api

import (
	"database/sql"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/AyKrimino/JobSeekerAPI/docs"
	"github.com/AyKrimino/JobSeekerAPI/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler(s.db)
	userHandler.RegisterRoutes(subrouter)

	// Swagger docs
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Listening on: ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
