package posts

import (
	"net/http"

	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi"
)

func newPutPostsHandler(postService PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement
		httpapi.RespondWithJSON(w, http.StatusOK, nil)
	}
}
