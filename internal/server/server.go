package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"users/config"
	postgresAdapter "users/internal/adapter/postgres"
	"users/internal/adapter/repository/oreshnik"
	"users/internal/infrastructure/database"
	"users/internal/infrastructure/database/postgres"
	"users/internal/ports/repository"
	"users/internal/usecase"
	"users/internal/usecase/users"
)

type Server struct {
	cfg *config.Config

	oreshnikDB *postgres.Postgres

	// repositories
	cartsRepository  repository.CartsRepository
	ordersRepository repository.OrdersRepository

	// services
	cartsUseCase  usecase.CartsUseCase
	ordersUseCase usecase.OrdersUseCase

	router *chi.Mux
	server *http.Server
}

func New(cfg *config.Config) (*Server, error) {
	s := &Server{
		cfg: cfg,
	}

	if err := s.init(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) init() error {
	if err := s.initDB(); err != nil {
		return fmt.Errorf("init db: %v", err)
	}
	if err := database.MigrateOreshnikDB(s.oreshnikDB); err != nil {
		return fmt.Errorf("migrate static db: %v", err)
	}

	s.initRepositories()
	s.initUseCases()
	s.initRouter()
	s.initHTTPServer()

	return nil
}

func (s *Server) initDB() error {
	var err error

	s.oreshnikDB, err = postgresAdapter.Connect(s.cfg.Server.StaticData.Connection)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) initRepositories() {
	s.cartsRepository = oreshnik.NewCartsRepository(s.oreshnikDB)
	s.ordersRepository = oreshnik.NewOrdersRepository(s.oreshnikDB)
}

func (s *Server) initUseCases() {
	s.cartsUseCase = users.NewCartsUseCase(s.cartsRepository)
	s.ordersUseCase = users.NewOrdersUseCase(s.ordersRepository)
}

func (s *Server) initHTTPServer() {
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.cfg.Server.Addr, s.cfg.Server.Port),
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (s *Server) Run() {
	log.Println("Server started")

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
