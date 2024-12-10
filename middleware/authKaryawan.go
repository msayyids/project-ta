package middleware

import (
	"context"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

type contextKeyKaryawan string

const KaryawannKey = contextKeyAdmin("karyawan")

type KaryawanContext struct {
	ID   int
	Role string
}

type AuthenticationKaryawan struct {
	UserService    service.UserServiceInj
	LayananService service.LayananServiceInj
}

func NewAuthKaryawan(us service.UserServiceInj, ls service.LayananServiceInj) AuthenticationAdmin {
	return AuthenticationAdmin{UserService: us, LayananService: ls}
}

func (a AuthenticationAdmin) AuthKaryawan(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("token")
		if token == "" {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "FORBIDDEN",
				Data:    "Access denied",
			}, http.StatusUnauthorized)
			return
		}

		claim, err := helper.ValidateToken(token)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "FORBIDDEN",
				Data:    "Access denied",
			}, http.StatusUnauthorized)
			return
		}

		id := int(claim["id"].(float64))
		role := claim["role"].(string)

		_, err = a.UserService.FindUserById(r.Context(), id)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "FORBIDDEN",
				Data:    "Access denied",
			}, http.StatusUnauthorized)
			return
		}
		_, err = a.UserService.FindUserByRole(r.Context(), role)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "FORBIDDEN",
				Data:    "Access denied",
			}, http.StatusUnauthorized)
			return
		}

		adminCtx := KaryawanContext{ID: id, Role: role}
		ctxAuth := context.WithValue(r.Context(), adminKey, adminCtx)
		r = r.WithContext(ctxAuth)

		next(w, r, ps)
	}
}
