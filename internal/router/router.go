package router

import (
	"database/sql"
	"net/http"
)

func New(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fs))
	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	// API маршруты
	RegisterSwagger(mux)
	registerArticleRoutes(mux, db)
	registerAuthRoutes(mux, db)
	registerCategoryRoutes(mux, db)
	registerUserRoutes(mux, db)
	return mux
}
