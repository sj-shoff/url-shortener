package delete

import (
	"net/http"

	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=URLDeleter
type URLDeleter interface {
	DeleteURL(alias string) error
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "id")
		if alias == "" {
			log.Error("empty alias")
			render.JSON(w, r, resp.Error("empty alias"))
			return
		}

		if err := urlDeleter.DeleteURL(alias); err != nil {
			log.Error("failed to delete URL", sl.Err(err), slog.String("alias", alias))
			render.JSON(w, r, resp.Error("internal server error"))
			return
		}

		responseDelete(w, r, alias)
	}

}

func responseDelete(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Alias:    alias,
	})
}
