package server

import (
	"users/internal/handlers/orders"

	"github.com/go-chi/chi/v5"

	"users/internal/handlers/carts"
)

func (s *Server) initRouter() {
	s.router = chi.NewRouter()

	s.router.Route("/api", func(r chi.Router) {
		r.Route("/cart", s.registerCartsRoutes)
		r.Route("/orders", s.registerOrdersRoutes)
	})
}

func (s *Server) registerCartsRoutes(r chi.Router) {
	r.Get("/", carts.GetCart(s.cartsUseCase))
	r.Patch("/", carts.UpdateCart(s.cartsUseCase))
}

func (s *Server) registerOrdersRoutes(r chi.Router) {
	r.Post("/", orders.CreateOrder(s.ordersUseCase))
	r.Patch("/", orders.UpdateOrder(s.ordersUseCase))
	r.Get("/", orders.GetOrders(s.ordersUseCase))
}
