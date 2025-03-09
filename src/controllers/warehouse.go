package controllers

import (
	"net/http"

	"warehouse-service/helpers"
)

func GetWarehousesListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetWarehousesList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/warehouses/list")
	}
}

func GetWarehouseV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetWarehouse(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/warehouses/get")
	}
}

func WarehouseBookItemV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.BookItem(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/warehouses/book-item")
	}
}

func WarehouseRollbackBookV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.RollbackBook(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/warehouses/rollback-book")
	}
}
