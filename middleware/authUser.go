package middleware

import (
	"context"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

type AuthenticationUser struct {
	UserService    service.UserServiceInj
	LayananService service.LayananServiceInj
}

type contextKeyUser string

const UserKey = contextKeyUser("users")

type UserContext struct {
	ID int
}

func NewAuthUser(us service.UserServiceInj, ls service.LayananServiceInj) AuthenticationUser {
	return AuthenticationUser{UserService: us, LayananService: ls}
}

func (a AuthenticationUser) AuthUser(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("token")
		if token == "" {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    nil,
			}, http.StatusUnauthorized)
			return
		}

		claim, err := helper.ValidateToken(token)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    nil,
			}, http.StatusUnauthorized)
			return
		}

		id := int(claim["id"].(float64))

		_, err = a.UserService.FindUserById(r.Context(), id)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    nil,
			}, http.StatusUnauthorized)
			return
		}
		userCtx := UserContext{ID: id}
		ctxAuth := context.WithValue(r.Context(), UserKey, userCtx)
		r = r.WithContext(ctxAuth)

		next(w, r, ps)
	}
}
