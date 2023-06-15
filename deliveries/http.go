package deliveries

import (
	"app/config"
	"app/controllers"
	"github.com/go-chi/chi"
	"net/http"
)

const (
	apiKeyHeader = "X-API-Key"
)

// APIKeyAuth middleware to check the API key in the request header
func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get(apiKeyHeader)

		if key != config.ServiceConfig.Server.ApiKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func StartHTTPServer(adsController *controllers.AdsController) {
	r := chi.NewRouter()

	r.Route("/v1", func(adsRouter chi.Router) {
		adsRouter.Use(APIKeyAuth)
		adsRouter.Get("/advertisement", adsController.GetUserAd)
	})

	http.ListenAndServe(":"+config.ServiceConfig.Server.Port, r)
}
