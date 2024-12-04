package orders

import (
	"encoding/json"
	"fmt"
	"net/http"

	"users/internal/models/dto"
	"users/internal/usecase"
	httpErr "users/pkg/http/error"
)

var statusEnum = map[string]struct{}{
	"В работе":     {},
	"Доставляется": {},
	"Получен":      {},
	"Отменен":      {},
	"Отказ":        {},
}

func UpdateOrder(uc usecase.OrdersUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order dto.UpdateOrder

		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			httpErr.InternalError(w, err)
			return
		}

		if _, ok := statusEnum[order.Status]; !ok {
			httpErr.BadRequest(w, fmt.Errorf("incorrect order status"))
			return
		}

		if err := uc.UpdateOrderStatus(r.Context(), order); err != nil {
			httpErr.InternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
