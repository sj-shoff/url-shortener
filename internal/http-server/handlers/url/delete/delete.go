package delete

import (
	"net/http"

	"log/slog"

	"github.com/go-chi/chi/v5"
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
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")

		if alias == "" {
			log.Info("empty alias")
			render.JSON(w, r, resp.Error("empty alias"))
			return
		}

		err := urlDeleter.DeleteURL(alias)
		if err != nil {
			log.Error("failed to delete url", sl.Err(err), slog.String("alias", alias))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		log.Info("url deleted", slog.String("alias", alias))

		render.JSON(w, r, resp.OK())
	}
}
