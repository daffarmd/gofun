package controller

import (
	"net/http"

	services "github.com/daffarmd/gofun/pzn/section-12/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
