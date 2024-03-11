package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"net/http"
	"path"
	"time"
)

func (s *Server) initializeHTTPServer() {

	r := chi.NewRouter()

	// Setup middlewares
	r.Use(middleware.Logger, middleware.Recoverer, middleware.RequestID, CSPMiddleware, cors.Handler(cors.Options{
		AllowedOrigins:   []string{constants.All},
		AllowedMethods:   []string{constants.Get, constants.Post, constants.Put, constants.Delete, constants.Patch, constants.Options},
		AllowedHeaders:   []string{constants.All},
		AllowCredentials: true,
	}), s.middleware.SecurityHeadersMiddleware, s.middleware.RateLimiterMiddleware())

	// Add Swagger UI only in development environment
	if s.config.Environment == constants.Development {
		swaggerDir := path.Join("doc", "openapi", "swagger")
		r.Mount("/swagger", http.StripPrefix("/swagger", http.FileServer(http.Dir(swaggerDir))))
		r.Get("/swagger/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, path.Join("docs", "openapi", "swagger.json"))
		})
		fmt.Printf("Swagger UI available at http://localhost:%d/swagger-ui\n", s.config.Server.HTTP.Port)
	}

	// HTTP server setup
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", s.config.Server.HTTP.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(s.config.Server.HTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.HTTP.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.config.Server.HTTP.IdleTimeout) * time.Second,
	}
}

func CSPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csp := "default-src 'self'; style-src 'self' 'unsafe-inline'; script-src 'self' 'unsafe-inline';"
		w.Header().Set("Content-Security-Policy", csp)
		next.ServeHTTP(w, r)
	})
}
