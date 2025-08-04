package routes

import (
	"net/http"

	"github.com/ravigill3969/Log-Aggregator-status-go/internal/handlers"
)

func RoutesForAggregation(mux *http.ServeMux) {
	handler := &handlers.AllRounderStruct{}
	mux.HandleFunc("POST /log/print/save", handler.StatusHandler)
}
