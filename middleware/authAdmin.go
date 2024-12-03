package middleware

import (
	"context"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

type contextKeyAdmin string

const adminKey = contextKeyAdmin("admin")

type AdminContext struct {
	ID   int
	Role string
}

type AuthenticationAdmin struct {
	UserService    service.UserServiceInj
	LayananService service.LayananServiceInj
}

func NewAuthAdmin(us service.UserServiceInj, ls service.LayananServiceInj) AuthenticationAdmin {
	return AuthenticationAdmin{UserService: us, LayananService: ls}
}

func (a AuthenticationAdmin) AuthAdmin(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

		_, err = a.UserService.FindUserById(r.Context(), id)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}
		_, err = a.UserService.FindUserByRole(r.Context(), role)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}

		adminCtx := KaryawanContext{ID: id, Role: role}
		ctxAuth := context.WithValue(r.Context(), adminKey, adminCtx)
		r = r.WithContext(ctxAuth)

		next(w, r, ps)
	}
}
