package middlewares

import (
	"net/http"

	"profile/api/auth"
	"profile/api/models"
	"profile/api/utils"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidToken(r)
		if err != nil {
			errObj := models.Error{}
			errObj.HasError(true, http.StatusUnauthorized, "Unauthorized")
			utils.ResponseJSON(w, http.StatusUnauthorized, "Token invalido", "", errObj)
			return
		}
		next(w, r)
	}
}
