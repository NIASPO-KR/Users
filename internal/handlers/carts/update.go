package carts

import (
	"encoding/json"
	"errors"
	"net/http"

	"users/internal/errs"
	"users/internal/models/dto"
	"users/internal/usecase"
	httpErr "users/pkg/http/error"
	"users/pkg/http/writer"
)

func UpdateCart(uc usecase.CartsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cartItem dto.CartItem

		if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
			httpErr.InternalError(w, err)
			return
		}

		err := uc.UpdateCartItem(r.Context(), dto.UpdateCartItem{
			CartItem: cartItem,
			UserID:   usecase.MockUsername,
		})
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				httpErr.NotFound(w, err)
				return
			}
			httpErr.BadRequest(w, err)
			return
		}

		writer.WriteStatusOK(w)
	}
}
