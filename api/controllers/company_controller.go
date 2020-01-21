package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/silas-ss/ms-company/api/models"
	"github.com/silas-ss/ms-company/api/helpers"
)

func (server *Server) CreateCompany(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	company := models.Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	company.Prepare()
	companyCreated, err := company.SaveCompany(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-type", "application/json")

	responses.JSON(w, http.StatusCreated, companyCreated)
}

func (server *Server) FindAllCompanies(w http.ResponseWriter, r *http.Request) {
	company := models.Company{}

	companies, err := company.FindAllCompanies(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-type", "application/json")

	responses.JSON(w, http.StatusOK, companies)
}

func (server *Server) FindCompanyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}

	c, err := company.FindCompanyByID(server.DB, companyID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-type", "application/json")

	responses.JSON(w, http.StatusOK, c)
}

func (server *Server) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	companyId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	err = server.DB.Debug().Model(models.Company{}).Where("id = ?", companyId).Take(&company).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Company not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	companyUpdate := models.Company{}
	err = json.Unmarshal(body, &companyUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	companyUpdate.Prepare()

	companyUpdate.ID = company.ID

	companyUpdated, err := companyUpdate.UpdateCompany(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	w.Header().Set("Content-type", "application/json")
	responses.JSON(w, http.StatusOK, companyUpdated)
}

func (server *Server) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	companyId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	company := models.Company{}
	err = server.DB.Debug().Model(models.Company{}).Where("id = ?", companyId).Take(&company).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = company.DeleteCompany(server.DB, companyId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	responses.JSON(w, http.StatusNoContent, "")
}
