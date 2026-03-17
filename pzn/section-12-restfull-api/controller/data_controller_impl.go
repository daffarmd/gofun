package controller

import (
	"net/http"
	"strconv"

	"github.com/daffarmd/gofun/helper"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/model/web"
	"github.com/daffarmd/gofun/pzn/section-12-restfull-api/service"
	"github.com/julienschmidt/httprouter"
)

type DataControllerImpl struct {
	DataService service.DataService
}

func NewDataController(sd service.DataService) DataController {
	return &DataControllerImpl{
		DataService: sd,
	}
}

func (t *DataControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	dataCreateRequest := web.DataCreateRequest{}
	helper.ReadFromReqBody(r, &dataCreateRequest)

	dataResponse := t.DataService.Create(r.Context(), dataCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   dataResponse,
	}

	helper.WriteResponseBody(w, webResponse)

}

func (t *DataControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	dataUpdateRequest := web.DataUpdateRequest{}
	helper.ReadFromReqBody(r, &dataUpdateRequest)
	defer r.Body.Close()
	dataId := params.ByName("dataId")
	dataIdConv, err := strconv.Atoi(dataId)
	helper.PanicIfErr(err)

	dataUpdateRequest.Id = dataIdConv

	dataResponse := t.DataService.Update(r.Context(), dataUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   dataResponse,
	}

	helper.WriteResponseBody(w, webResponse)

}

func (t *DataControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dataId := params.ByName("dataId")
	dataIdConv, err := strconv.Atoi(dataId)
	helper.PanicIfErr(err)

	t.DataService.Delete(r.Context(), dataIdConv)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
	}

	helper.WriteResponseBody(w, webResponse)
}

func (t *DataControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dataId := params.ByName("dataId")
	dataIdConv, err := strconv.Atoi(dataId)
	helper.PanicIfErr(err)

	dataResponse := t.DataService.FindById(r.Context(), dataIdConv)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   dataResponse,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (t *DataControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sort := r.URL.Query().Get("sort")

	dataResponse := t.DataService.FindAll(r.Context())
	if sort == "asc" {
		dataResponse = t.DataService.FindAllAsc(r.Context())
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   dataResponse,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (t *DataControllerImpl) FindAllAsc(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dataResponse := t.DataService.FindAllAsc(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   dataResponse,
	}

	helper.WriteResponseBody(w, webResponse)
}
