package handler

import (
	"errors"
	"main/core/controller"
	"main/core/database"
	"main/core/helper"
	"main/core/model"
	"main/core/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	Interface
	Controller controller.Interface
	Rules      helper.Interface
}

func NewHandler(repository database.Interface) Interface {
	return &Handler{
		Controller: controller.NewController(repository),
		Rules:      helper.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "id") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	response, err := h.Controller.ListOne(id)
	if err != nil {
		utils.StatusInternalServerError(w, r, err)
		return
	}

	utils.StatusOK(w, r, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		utils.StatusInternalServerError(w, r, err)
		return
	}

	utils.StatusOK(w, r, response)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.getBody(r, "")
	if err != nil {
		utils.StatusBadRequest(w, r, err)
		return
	}

	id, err := h.Controller.Create(productBody)
	if err != nil {
		utils.StatusInternalServerError(w, r, err)
	}

	utils.StatusOK(w, r, id)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	productBody, err := h.getBody(r, id)
	if err != nil {
		utils.StatusBadRequest(w, r, err)
		return
	}

	err = h.Controller.Update(id, productBody)
	if err != nil {
		utils.StatusInternalServerError(w, r, err)
	}

	utils.StatusNoContent(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.Controller.Remove(id)
	if err != nil {
		utils.StatusInternalServerError(w, r, err)
		return
	}

	utils.StatusNoContent(w, r)
}

func (h *Handler) getBody(r *http.Request, id string) (*model.Product, error) {
	productBody := &model.Product{}
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)
	if err != nil {
		return &model.Product{}, errors.New("body is required")
	}

	productParsed, err := model.InterfaceToModel(body)
	if err != nil {
		return &model.Product{}, errors.New("error on converting body to model")
	}

	setDefaultValues(&productParsed, id)
	return &productParsed, nil
}

func setDefaultValues(product *model.Product, id string) {
	product.UpdatedAt = time.Now()
	if id == "" {
		product.ID = uuid.NewString()
		product.CreatedAt = time.Now()
	} else {
		product.ID = id
	}
}
