package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/utils"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		token, err := utils.VerifyToken(r)

		if err != nil {
			resErr := responses.ErrorResponseModel{
				FieldName: "",
				Message: err.Error(),
			}

			errors := responses.NewErrorResponse(resErr)

			out, _ := json.Marshal(errors)

			http.Error(rw, string(out), http.StatusBadRequest)
			return
		}

		if !token.Valid {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	})
}