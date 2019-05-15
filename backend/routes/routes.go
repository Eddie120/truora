package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"truora/backend/app/Models"
	"truora/backend/app/api"
)


var arreglo_endpoints = Models.Rutas{

	Models.Ruta{"/", "GET", api.Index},
	Models.Ruta{"/llave", "POST", api.CrearLlave},
	Models.Ruta{"/llaves", "GET", api.ListarLlaves},
	Models.Ruta{"/llave/encriptar", "POST",  api.EncriptarMensaje},
	Models.Ruta{"/llave/desencriptar", "POST",  api.DesencriptarMensaje},
}

func CargarRutas() *chi.Mux {
	r := chi.NewRouter()

	// habilitamos los CORS del servidor
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		// AllowOriginFunc:  func(r *api.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		//AllowCredentials: true,
		//MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	if len(arreglo_endpoints) > 0 {
		for i := 0; i < len(arreglo_endpoints); i ++ {
			r.MethodFunc(arreglo_endpoints[i].Metodo, arreglo_endpoints[i].Patron, arreglo_endpoints[i].Funcion)
		}
	}

	return r
}