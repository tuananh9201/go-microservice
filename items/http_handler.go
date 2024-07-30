package main

import (
	"context"
	"net/http"

	"log"

	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tuananh9201/commons"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{db}
}

func (h *handler) Register(r *mux.Router) {
	r.HandleFunc("/items", CreateItemHanlder(h.db)).Methods(http.MethodPost)
	r.HandleFunc("/items", GetItemsHandler(h.db)).Methods(http.MethodGet)
	r.HandleFunc("/items/{id}", GetItemByIdHandler(h.db)).Methods(http.MethodGet)
	r.HandleFunc("/items/{id}", UpdateItemHandler(h.db)).Methods(http.MethodPut)
	r.HandleFunc("/items/{id}", DeleteItemHandler(h.db)).Methods(http.MethodDelete)
}

func CreateItemHanlder(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data ItemCreate
		if err := common.ReadJSON(r, &data); err != nil {
			common.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		log.Println("Create item", data)
		store := NewSQLStore(db)
		uc := NewItemUsecase(store)
		if err := uc.CreateNewItem(context.Background(), &data); err != nil {
			common.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}

		common.WriteJSON(w, http.StatusCreated, common.SimpleSuccessResponse(data))
	}
}

func GetItemsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paging := ParsePaging(r.URL.Query())
		filter := ParseFilter(r.URL.Query())
		store := NewSQLStore(db)
		uc := NewItemUsecase(store)
		items, p, err := uc.GetItems(context.Background(), &paging, &filter)
		if err != nil {
			common.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		common.WriteJSON(w, http.StatusOK, common.NewSuccessResponse(items, p, filter))
	}
}

func ParsePaging(query url.Values) common.Paging {
	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit < 1 {
		limit = 10
	}
	return common.Paging{Page: page, Limit: limit}
}

// ParseFilter extracts the filter parameters from the query string
func ParseFilter(query url.Values) ItemFilter {
	return ItemFilter{Name: query.Get("name")}
}

func GetItemByIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		store := NewSQLStore(db)
		uc := NewItemUsecase(store)
		item, err := uc.GetItemByID(context.Background(), id)
		if err != nil {
			common.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		common.WriteJSON(w, http.StatusOK, common.SimpleSuccessResponse(item))
	}
}

func UpdateItemHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		var data ItemCreate
		if err := common.ReadJSON(r, &data); err != nil {
			common.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		store := NewSQLStore(db)
		uc := NewItemUsecase(store)
		if err := uc.UpdateItem(context.Background(), id, &data); err != nil {
			common.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		common.WriteJSON(w, http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func DeleteItemHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		store := NewSQLStore(db)
		uc := NewItemUsecase(store)
		if err := uc.DeleteItem(context.Background(), id); err != nil {
			common.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		common.WriteJSON(w, http.StatusOK, common.SimpleSuccessResponse("Deleted"))
	}
}
