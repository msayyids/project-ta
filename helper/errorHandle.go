package helper

import (
	"fmt"
	"net/http"
	"project-ta/entity"
	"runtime/debug"
)

func ErrorHandler(w http.ResponseWriter, statuscode int, message string) {
	switch statuscode {
	case http.StatusBadRequest: // 400 Bad Request
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    message,
		}, http.StatusBadRequest)

	case http.StatusUnauthorized: // 401 Unauthorized
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: "UNAUTHORIZED",
			Data:    message,
		}, http.StatusUnauthorized)

	case http.StatusForbidden: // 403 Forbidden
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusForbidden,
			Message: "FORBIDDEN",
			Data:    message,
		}, http.StatusForbidden)

	case http.StatusNotFound: // 404 Not Found
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    message,
		}, http.StatusNotFound)

	case http.StatusInternalServerError: // 500 Internal Server Error
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Data:    message,
		}, http.StatusInternalServerError)

	case http.StatusBadGateway: // 502 Bad Gateway
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadGateway,
			Message: "BAD GATEWAY",
			Data:    message,
		}, http.StatusBadGateway)

	case http.StatusServiceUnavailable: // 503 Service Unavailable
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "SERVICE UNAVAILABLE",
			Data:    message,
		}, http.StatusServiceUnavailable)

	default:
		ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Data:    fmt.Sprintf("An unexpected error occurred: %s", message),
		}, http.StatusInternalServerError)
	}
}

// PanicHandlerWrapper adalah wrapper untuk menangani panic dan memanggil ErrorHandler
func PanicHandlerWrapper(w http.ResponseWriter, _ *http.Request, err interface{}) {
	// utntuk debug
	fmt.Printf("PANIC: %v\n", err)
	fmt.Printf("Stack Trace:\n%s\n", debug.Stack())

	ErrorHandler(w, http.StatusInternalServerError, fmt.Sprintf("An unexpected error occurred: %v", err))
}
