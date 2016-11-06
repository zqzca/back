package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pressly/chi"
)

// Paginate adds the values "per_page" and "page" to the context with values
// from the query params.
func Paginate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rawPerPage := chi.URLParam(r, "per_page")
		rawPage := chi.URLParam(r, "page")

		if len(rawPerPage) == 0 {
			rawPerPage = "20"
		}

		if len(rawPage) == 0 {
			rawPerPage = "0"
		}

		perPage, _ := strconv.Atoi(rawPerPage)
		page, _ := strconv.Atoi(rawPage)

		ctx := r.Context()
		ctx = context.WithValue(ctx, 1001, perPage)
		ctx = context.WithValue(ctx, 1002, page)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
