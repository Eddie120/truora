package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"truora/backend/app/models"
	"truora/backend/app/api"
)


var EndPoints = models.Routes{

	models.Route{"/key", "POST", api.CreateKey},
	models.Route{"/keys", "GET", api.Index},
	models.Route{"/key/encrypt", "POST",  api.Encrypt},
	models.Route{"/key/decrypt", "POST",  api.Decrypt},
}

func LoadRoutes() *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options {
		AllowedOrigins:   []string{"http://localhost:8081/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	r.Use(cors.Handler)

	for i := 0; i < len(EndPoints); i ++ {
		r.MethodFunc(EndPoints[i].Method, EndPoints[i].Pattern, EndPoints[i].Function)
	}


	return r
}