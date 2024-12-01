package server

import (
	"github.com/go-chi/chi/v5"

	"users/internal/handlers/carts"
)

func (s *Server) initRouter() {
	s.router = chi.NewRouter()

	s.router.Route("/api", func(r chi.Router) {
		r.Route("/cart", s.registerCartsRoutes)
	})
}

func (s *Server) registerCartsRoutes(r chi.Router) {
	r.Get("/", carts.GetCart(s.cartsUseCase))
	r.Patch("/", carts.UpdateCart(s.cartsUseCase))
}
