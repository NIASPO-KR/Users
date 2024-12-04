package orders

import (
	"net/http"

	"users/internal/usecase"
	httpErr "users/pkg/http/error"
	"users/pkg/http/writer"
)

func GetOrders(uc usecase.OrdersUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := uc.GetOrders(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, orders)
	}
}
