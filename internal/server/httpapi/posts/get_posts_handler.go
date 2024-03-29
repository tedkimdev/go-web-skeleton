package posts

import (
	"net/http"

	"github.com/tedkimdev/go-web-skeleton/api/openapi"
	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi"
)

func newGetPostsHandler(postService PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := postService.GetPosts(r.Context())
		if err != nil {
			httpapi.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		res := openapi.ModelPostsToDTO(posts)
		httpapi.RespondWithJSON(w, http.StatusOK, res)
	}
}
