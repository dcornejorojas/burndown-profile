package middlewares

import (
	"net/http"
	"profile/api/auth"
	"profile/api/models"
	"profile/api/utils"
)

func SetAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidToken(r)
		if err != nil {
			errObj := models.Error{}
			errObj.HasError(true, http.StatusUnauthorized, "Unauthorized")
			utils.ResponseJSON(w, http.StatusUnauthorized, "Token invalido", []string{}, errObj)
			return
		}
		next(w, r)
	}
}
