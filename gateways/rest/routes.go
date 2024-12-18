package rest

import (
	"github.com/go-chi/chi/v5"
)

func registerRoutes(handler *handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Post("/", handler.createProduct)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getProduct)
		})
	})

	return r
}
