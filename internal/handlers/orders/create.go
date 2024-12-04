package orders

import (
	"encoding/json"
	"net/http"

	"users/internal/models/dto"
	"users/internal/usecase"
	httpErr "users/pkg/http/error"
	"users/pkg/http/writer"
)

func CreateOrder(uc usecase.OrdersUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order dto.CreateOrder

		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			httpErr.InternalError(w, err)
			return
		}
		orderID, err := uc.CreateOrder(r.Context(), order)
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, dto.CreatedIDInt{ID: orderID})
	}
}
