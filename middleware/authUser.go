package middleware

import (
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

func NewAuthUser(us service.UserServiceInj, ls service.LayananServiceInj) AuthenticationUser {
	return AuthenticationUser{UserService: us, LayananService: ls}
}

func (a AuthenticationUser) AuthUser(next httprouter.Handle) httprouter.Handle {
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

		_, err = a.UserService.FindUserById(r.Context(), id)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}

		next(w, r, ps)
	}
}
