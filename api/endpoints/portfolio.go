package endpoints

import (
	"fmt"
	"github.com/paulo-granthon/portfolio_apis/server"
	"github.com/paulo-granthon/portfolio_apis/service"
	"net/http"
	"strconv"

	"github.com/ztrue/tracerr"
)

func PortfolioEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/portfolio/{id}",
			Methods: []server.Method{
				server.NewMethod("GET", GetPortfolio),
			},
		},
		{
			Path: "/portfolio/{id}/markdown",
			Methods: []server.Method{
				server.NewMethod("GET", GetPortfolioMarkdown),
			},
		},
	}
}

// parsePortfolioUserId reads and validates the {id} path parameter.
func parsePortfolioUserId(w http.ResponseWriter, r *http.Request) (uint64, bool) {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		server.SendError(w, err, http.StatusBadRequest, "Parameter id not found")
		return 0, false
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		server.SendError(w, tracerr.Wrap(err), http.StatusBadRequest, "Parameter id is not a valid number")
		return 0, false
	}

	return id, true
}

// GetPortfolio godoc
// @Summary get the composed portfolio document of a user
// @Tags    portfolio
// @Produce json
// @Param   id   path     int  true  "user id"
// @Success 200  {object}  models.Portfolio
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /portfolio/{id} [get]
func GetPortfolio(s server.Server, w http.ResponseWriter, r *http.Request) error {
	id, ok := parsePortfolioUserId(w, r)
	if !ok {
		return nil
	}

	portfolio, err := s.Service.PortfolioService.Build(id)
	if err != nil {
		return server.SendError(w, err, http.StatusInternalServerError, "error building portfolio")
	}

	return server.WriteJSON(w, http.StatusOK, portfolio)
}

// GetPortfolioMarkdown godoc
// @Summary download the portfolio document of a user as markdown
// @Tags    portfolio
// @Produce text/markdown
// @Param   id   path     int  true  "user id"
// @Success 200  {string}  string
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /portfolio/{id}/markdown [get]
func GetPortfolioMarkdown(s server.Server, w http.ResponseWriter, r *http.Request) error {
	id, ok := parsePortfolioUserId(w, r)
	if !ok {
		return nil
	}

	portfolio, err := s.Service.PortfolioService.Build(id)
	if err != nil {
		return server.SendError(w, err, http.StatusInternalServerError, "error building portfolio")
	}

	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="portfolio-%d.md"`, id))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(service.RenderMarkdown(*portfolio)))
	return err
}
