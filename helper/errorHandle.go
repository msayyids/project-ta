package helper

import (
	"net/http"
	"project-ta/entity"

	"github.com/go-playground/validator/v10"
)

type NotFoundError struct {
	Error string
}

func NewAnotherError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}
	if validationErrors(writer, request, err) {
		return
	}

	// Default to internal server error if no other errors match
	internalServerError(writer, request, err)
}

// Handler for validation errors
func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := entity.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		ResponseBody(writer, webResponse)
		return true
	}
	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := entity.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		ResponseBody(writer, webResponse)
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	var errorString string
	if e, ok := err.(error); ok {
		errorString = e.Error()
	} else {
		errorString = "An unexpected error occurred"
	}

	webResponse := entity.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   errorString,
	}

	ResponseBody(writer, webResponse)
}
