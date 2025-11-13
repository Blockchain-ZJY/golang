package handler

import (
	respondjson "golang/rssagg/RespondJSON"
	"net/http"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	respondjson.RespondWithError(w, 400, "Something went wrong")
}
