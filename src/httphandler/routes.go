package httphandler

import (
	"net/http"
	"regexp"

	"warehouse-service/controllers"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// notifications
	newRoute(http.MethodGet, "/v1/warehouses/list", controllers.GetWarehousesListV1),
	newRoute(http.MethodGet, "/v1/warehouses/get/([0-9]+)", controllers.GetWarehouseV1),
	newRoute(http.MethodPost, "/v1/warehouses/book-item", controllers.WarehouseBookItemV1),
	newRoute(http.MethodPost, "/v1/warehouses/rollback-book", controllers.WarehouseRollbackBookV1),
}
