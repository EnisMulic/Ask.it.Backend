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

			rw.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(rw).Encode(errors)

			return
		}

		if !token.Valid {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		next.ServeHTTP(rw, r)
	})
}