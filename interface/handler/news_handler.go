package handler

import (
	"net/http"
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type NewsHandler interface {
	Get() echo.HandlerFunc
	Mock() echo.HandlerFunc
}

type newsHandler struct {
	newsUsecase usecase.NewsUsecase
}

func NewNewsHandler(newsUsecase usecase.NewsUsecase) NewsHandler {
	return &newsHandler{newsUsecase: newsUsecase}
}

type requestNews struct{}

type responseNews struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
	Tag   string `json:"tag"`
	Link  string `json:"link"`
}

func (h *newsHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("uuid")
		if uuid != config.GetUser(c).UUID {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid uuid")
		}

		user := config.GetUser(c)
		getNews, err := h.newsUsecase.Get(&domain.User{
			ID:       user.ID,
			Password: user.Password,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseNews{}
		for _, news := range getNews {
			res = append(res, responseNews{
				ID:    news.ID,
				Title: news.Title,
				Date:  news.Date,
				Tag:   news.Tag,
				Link:  news.Link,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (h *newsHandler) Mock() echo.HandlerFunc {
	return func(c echo.Context) error {
		raw, err := os.ReadFile("mocks/data/news.json")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseNews
		if err := json.Unmarshal(raw, &res); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}