package router

import (
	"net/http"

	"github.com/desa-wisata/desa-wisata-server/delivery/handler"
	"github.com/desa-wisata/library/middleware"
	"github.com/go-chi/chi"
)

type route struct {
	api handler.Base
}

//Route ...
type Route interface {
	Mux() http.Handler
}

//InitRoute ...
func InitRoute(api handler.Base) Route {
	return &route{api}
}

func (c *route) Mux() http.Handler {

	router := chi.NewRouter()
	router.Use(middleware.LogMiddleware)
	router.Group(func(r chi.Router) {
		// Destination
		r.Get("/destination", c.api.GetDestination)
	})

	return router
}
