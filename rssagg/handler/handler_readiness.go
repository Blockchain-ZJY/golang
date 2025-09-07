package handler

import (
	respondjson "golang/rssagg/RespondJSON"
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondjson.RespondWithJSON(w, 200, struct {
		Info string `json:"info"`
	}{
		Info: "everything is ok",
	})
}
