package middleware

import (
	"cloud-disk/core/helper"
	"net/http"
	"strconv"
	"strings"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		if auth==""{
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: token is null"))
			return
		}

		if strings.Index(auth,"Bearer")==-1{
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: token formate is not correct"))
			return
		}

		token := strings.Replace(auth, "Bearer ","",1)

		uc, err := helper.ParseToken(token)
		if err!=nil{
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized:"+err.Error()))
			return
		}
		r.Header.Set("UserId",strconv.Itoa(uc.Id))
		r.Header.Set("UserIdentity",uc.Identity)
		r.Header.Set("UserName",uc.Name)
		next(w, r)
	}
}
