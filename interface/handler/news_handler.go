package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type NewsHandler interface {
	Get() echo.HandlerFunc
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
		getNews, err := h.newsUsecase.Get(&domain.User{
			Id:     os.Getenv("TEST_ID"),
			Passwd: os.Getenv("TEST_PW"),
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseNews{}
		for _, news := range getNews {
			res = append(res, responseNews{
				ID:    news.Id,
				Title: news.Title,
				Date:  news.Date,
				Tag:   news.Tag,
				Link:  news.Link,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}
