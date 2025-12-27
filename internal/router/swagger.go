package router

import (
	"net/http"
	_ "newww/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwagger(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
