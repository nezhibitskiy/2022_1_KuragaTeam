package delivery

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/moviesCompilations"
	"net/http"
	"strconv"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	MCByPersonURL = "api/v1/movieCompilations/person/:person_id"
	MCByMovieURL  = "api/v1/movieCompilations/movie/:movie_id"
	MCByGenreURL  = "api/v1/movieCompilations/genre/:genre_id"
	MCTopURL      = "api/v1/movieCompilations/top"
	MCYearTopURL  = "api/v1/movieCompilations/yearTop/:year"
	MCDefaultURL  = "/api/v1/movieCompilations"
)

type handler struct {
	movieCompilationsService moviesCompilations.Service
}

func NewHandler(service moviesCompilations.Service) *handler {
	return &handler{movieCompilationsService: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(MCDefaultURL, h.GetMoviesCompilations())
	router.GET(MCByMovieURL, h.GetMCByMovieID())
	router.GET(MCByGenreURL, h.GetMCByGenre())
	router.GET(MCByPersonURL, h.GetMCByPersonID())
	router.GET(MCTopURL, h.GetTopMC())
	router.GET(MCYearTopURL, h.GetYearTopMC())
}

func (h *handler) GetMoviesCompilations() echo.HandlerFunc {
	return func(context echo.Context) error {
		mainMoviesCompilations, err := h.movieCompilationsService.GetMainCompilations()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &mainMoviesCompilations)
	}
}

func (h *handler) GetMCByMovieID() echo.HandlerFunc {
	return func(context echo.Context) error {
		movieID, err := strconv.Atoi(context.Param("movie_id"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.movieCompilationsService.GetByMovie(movieID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedMC)
	}
}

func (h *handler) GetMCByGenre() echo.HandlerFunc {
	return func(context echo.Context) error {
		genreID, err := strconv.Atoi(context.Param("genre_id"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.movieCompilationsService.GetByGenre(genreID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedMC)
	}
}

func (h *handler) GetMCByPersonID() echo.HandlerFunc {
	return func(context echo.Context) error {
		personID, err := strconv.Atoi(context.Param("person_id"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.movieCompilationsService.GetByPerson(personID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedMC)
	}
}

func (h *handler) GetTopMC() echo.HandlerFunc {
	return func(context echo.Context) error {
		var limit int
		echo.QueryParamsBinder(context).Int("limit", &limit)
		if limit > 12 {
			limit = 12
		}
		selectedMC, err := h.movieCompilationsService.GetTop(limit)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedMC)
	}
}

func (h *handler) GetYearTopMC() echo.HandlerFunc {
	return func(context echo.Context) error {
		year, err := strconv.Atoi(context.Param("year"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.movieCompilationsService.GetTopByYear(year)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedMC)
	}
}
