package middleware

import (
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

type AuthenticationUser struct {
	UserService service.UserServiceInj
}

func NewAuthUser(us service.UserServiceInj) AuthenticationAdmin {
	return AuthenticationAdmin{UserService: us}
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

		_, err = a.UserService.FindUSerById(r.Context(), id)
		if err != nil {
			helper.ResponseBody(w, entity.WebResponse{
				Code:   401,
				Status: "UNAUTHORIZED",
				Data:   nil,
			})
			return
		}
	}
}
