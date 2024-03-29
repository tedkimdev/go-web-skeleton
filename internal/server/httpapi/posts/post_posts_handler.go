package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tedkimdev/go-web-skeleton/api/openapi"
	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi"
)

func newPostPostsHandler(postService PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req openapi.CreatePostRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			httpapi.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("marshalling request: %w", err))
			return
		}

		post, err := postService.CreatePost(r.Context(), req.Title, req.Body)
		if err != nil {
			httpapi.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("create post: %w", err))
			return
		}
		httpapi.RespondWithJSON(w, http.StatusOK, openapi.ModelPostToDTO(post))
	}
}
