package middleware

import (
	"context"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
)

type contextKeyAdmin string

const adminKey = contextKeyAdmin("admin")

type AdminContext struct {
	ID   int
	Role string
}

type AuthAdmininj interface {
	AuthAdmin(next http.Handler) http.Handler
}

type AuthenticationAdmin struct {
	UserService service.UserServiceInj
}

func NewAuthUser(us service.UserServiceInj) AuthenticationAdmin {
	return AuthenticationAdmin{UserService: us}
}

func (a AuthenticationUser) AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}

		claim, err := helper.ValidateToken(token)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}

		id := int(claim["id"].(float64))
		role := claim["role"].(string)

		_, err = a.UserService.FindUSerById(r.Context(), id)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}
		_, err = a.UserService.FindUSerByRole(r.Context(), role)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}

		adminCtx := UserContext{ID: id, Role: role}
		ctxAuth := context.WithValue(r.Context(), adminKey, adminCtx)
		r = r.WithContext(ctxAuth)

		next.ServeHTTP(w, r)
	})
}
