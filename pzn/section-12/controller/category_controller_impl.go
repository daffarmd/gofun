package controller

import (
	"encoding/json"
	"net/http"

	"github.com/daffarmd/gofun/pzn/section-12/helper"
	"github.com/daffarmd/gofun/pzn/section-12/model/web"
	services "github.com/daffarmd/gofun/pzn/section-12/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Save(ctx context.Context, request web.CategoryCreateRequest)
	webResponse := web.CategoryCreateRequest{} 

	panic("not implemented")
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
