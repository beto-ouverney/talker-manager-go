package handler

import (
	"fmt"
	"net/http"
)

// URLParam extracts a parameter from the URL by name
func URLParam(r *http.Request, name string) string {
	ctx := r.Context()
	params := ctx.Value("params").(map[string]string)
	fmt.Printf("%v", params)
	fmt.Println("PARANS")
	fmt.Println(params)
	return params[name]
}
