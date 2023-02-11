package handler

import (
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type NewsOnlyHandler interface {
	Get() echo.HandlerFunc
	Mock() echo.HandlerFunc
}

type newsOnlyHandler struct {
	newsOnlyUsecase usecase.NewsOnlyUsecase
}

func NewNewsOnlyHandler(newsOnlyUsecase usecase.NewsOnlyUsecase) NewsOnlyHandler {
	return &newsOnlyHandler{newsOnlyUsecase: newsOnlyUsecase}
}

type requestNewsOnly struct{}

type responseNewsOnly struct {
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Date       string   `json:"date"`
	Tag        string   `json:"tag"`
	Attachment []string `json:"attachment"`
}

func (h *newsOnlyHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("uuid")
		if uuid != config.GetUser(c).UUID {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid uuid")
		}

		id := c.Param("id")
		user := config.GetUser(c)
		getNewsOnly, err := h.newsOnlyUsecase.Get(
			id,
			&domain.User{
				ID:       user.ID,
				Password: user.Password,
			},
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseNewsOnly{
			{
				Title:      getNewsOnly.Title,
				Body:       getNewsOnly.Body,
				Date:       getNewsOnly.Date,
				Tag:        getNewsOnly.Tag,
				Attachment: getNewsOnly.Attachment,
			},
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (h *newsOnlyHandler) Mock() echo.HandlerFunc {
	return func(c echo.Context) error {
		raw, err := ioutil.ReadFile("mocks/data/news-only.json")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseNewsOnly
		if err := json.Unmarshal(raw, &res); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}