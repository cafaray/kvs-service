package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cafaray/kvs-service/internal/middleware"
	"github.com/cafaray/kvs-service/pkg/claim"
	"github.com/cafaray/kvs-service/pkg/response"
	"github.com/cafaray/kvs-service/pkg/user"
	"github.com/go-chi/chi"
)

type UserRouter struct {
	Repository user.Repository
}

func (ur *UserRouter) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/login/", ur.LoginHandler)
	r.
		With(middleware.Authorizator). // <-- here we are protecting Get operation.
		Get("/", ur.GetAllHandler)
	r.Post("/", ur.CreateHandler)
	r.Put("/{id}", ur.UpdateHandler)
	r.Delete("/{id}", ur.DeleteHandler)
	r.Get("/{id}", ur.GetOneHandler)
	return r
}

func (ur *UserRouter) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	storedUser, err := ur.Repository.GetByUsername(ctx, u.Username)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	if !storedUser.PasswordMatch(u.Password) {
		response.HTTPError(w, r, http.StatusUnauthorized, "Password does not match")
		return
	}
	c := claim.Claim{ID: int(storedUser.ID)}
	token, err := c.GetToken(os.Getenv("SIGNING_STRING"))
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"token": token})

}

func (ur *UserRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u) // <= here it's important that variable `u` should be a pointer, because we're  going to change its value later
	if err != nil {
		fmt.Println(err.Error())
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &u) // <= send the pointer to variable `u`
	if err != nil {
		fmt.Println(err.Error())
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	u.Password = ""
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), u.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"user": u})
	return
}
func (ur *UserRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"users": users})
}
func (ur *UserRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	id_string := chi.URLParam(r, "id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	u, err := ur.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"user": u})
}
func (ur *UserRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id_string := chi.URLParam(r, "id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var u user.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Update(ctx, uint(id), &u)
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}
func (ur *UserRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = ur.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}
