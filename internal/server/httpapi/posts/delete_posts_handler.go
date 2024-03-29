package posts

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi"
)

func newDeletePostsHandler(postService PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			httpapi.RespondWithError(w, http.StatusBadRequest, errors.New("id is empty"))
			return
		}

		err := postService.DeletePost(r.Context(), id)
		if err != nil {
			httpapi.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		httpapi.RespondWithJSON(w, http.StatusOK, nil)
	}
}
