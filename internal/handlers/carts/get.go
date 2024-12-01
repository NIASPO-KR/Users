package carts

import (
	"net/http"

	"users/internal/usecase"
	httpErr "users/pkg/http/error"
	"users/pkg/http/writer"
)

func GetCart(uc usecase.CartsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cart, err := uc.GetCarts(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, cart)
	}
}
