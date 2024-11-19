package middleware

import (
	"context"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
)

type contextKey string

const userKey = contextKey("user")

type UserContext struct {
	ID   int
	Role string
}

type AuthUsers interface {
	AuthUser(next http.Handler) http.Handler
}

type AuthenticationUser struct {
	UserService service.UserServiceInj
}

func NewAuthAdmin(us service.UserServiceInj) AuthenticationUser {
	return AuthenticationUser{UserService: us}
}

func (a AuthenticationUser) AuthUser(next http.Handler) http.Handler {
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

		// Retrieve the ID and role from token claims
		id := int(claim["id"].(float64))
		role := claim["role"].(string)

		// Verify ID and Role in the database
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

		userCtx := UserContext{ID: id, Role: role}
		ctxAuth := context.WithValue(r.Context(), userKey, userCtx)
		r = r.WithContext(ctxAuth)

		next.ServeHTTP(w, r)
	})
}
