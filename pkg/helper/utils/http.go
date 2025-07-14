package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetUrl(r *http.Request, key string) (int64, error) {

	var err error
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return 0, fmt.Errorf("parameter not valid")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parameter not valid")
	}

	return id, nil
}
