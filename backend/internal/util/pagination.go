package util

import (
    "net/http"
    "strconv"
)

func Pagination(r *http.Request) (int32, int32) {
    limit := int32(50)
    page := int32(1)
    if v, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil && v > 0 && v <= 100 { limit = int32(v) }
    if v, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && v > 0 { page = int32(v) }
    return limit, (page - 1) * limit
}
