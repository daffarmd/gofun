package exception

import (
	"net/http"

	"github.com/daffarmd/gofun/helper"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/web"
	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if notFoundError(w, r, err) {
		return
	}

	if badRequest(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteResponseBody(w, webResponse)
		return true
	} else {
		return false
	}

}

func badRequest(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Error Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteResponseBody(w, webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteResponseBody(w, webResponse)
}
