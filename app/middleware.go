package app

import (
	"context"
	"fmt"
	"net/http"
)

type MyPaginationKey string

const PaginationKey MyPaginationKey = "pagination"

const offsetDefault = 0
const limitDefault = 20
const pageDefault = 1
const sizeDefault = 20

func PaginationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		offset := query.Get("offset")
		limit := query.Get("limit")
		page := query.Get("page")
		size := query.Get("size")

		offsetValue, err := GetIntArg("offset", offset, offsetDefault)
		if err != nil {
			RespondWithError(w, 400, err.Error(), fmt.Sprintf("can not parse offset=`%s` to an integer", offset))
			return
		}

		limitValue, err := GetIntArg("limit", limit, limitDefault)
		if err != nil {
			RespondWithError(w, 400, err.Error(), fmt.Sprintf("can not parse limit=`%s` to an integer", limit))
			return
		}

		pageValue, err := GetIntArg("page", page, pageDefault)
		if err != nil {
			RespondWithError(w, 400, err.Error(), fmt.Sprintf("can not parse page=`%s` to an integer", page))
			return
		}

		sizeValue, err := GetIntArg("size", size, sizeDefault)
		if err != nil {
			RespondWithError(w, 400, err.Error(), fmt.Sprintf("can not parse size=`%s` to an integer", size))
			return
		}

		paginatioDTO := PaginationDTO{Page: pageValue, Limit: limitValue, Offset: offsetValue, Size: sizeValue}

		ctx := context.WithValue(r.Context(), PaginationKey, &paginatioDTO)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
