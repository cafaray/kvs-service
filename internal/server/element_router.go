package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cafaray/kvs-service/pkg/element"
	"github.com/cafaray/kvs-service/pkg/response"
	"github.com/go-chi/chi"
)

type ElementRouter struct {
	Repository element.Repository
}

func (er *ElementRouter) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", er.GetAllHandler)
	r.Post("/", er.CreateHandler)
	// r.Put("/{id}", ur.UpdateHandler)
	// r.Delete("/{id}", ur.DeleteHandler)
	r.Get("/{id}", er.GetOneHandler)
	return r
}

func (er *ElementRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var e element.Element
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = er.Repository.Create(ctx, &e)
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), e.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"element": e})
	return
}

func (er *ElementRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	elements, err := er.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"elements": elements})
	return
}

func (er *ElementRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	id_string := chi.URLParam(r, "id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	e, err := er.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"element:": e})
	return
}
