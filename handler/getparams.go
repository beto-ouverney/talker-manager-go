package handler

import (
	"net/http"
)

// URLParam extracts a parameter from the URL by name
func URLParam(r *http.Request, name string) string {
	ctx := r.Context()
	params := ctx.Value("params").(map[string]string)
	return params[name]
}
