// forum/utils/swagger.go

package utils

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitSwagger initializes Swagger middleware for Gorilla Mux framework
func InitSwagger(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
