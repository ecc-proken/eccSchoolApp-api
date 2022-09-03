package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type CalendarHandler interface {
	Get() echo.HandlerFunc
}

type calendarHandler struct {
	calendarUsecase usecase.CalendarUsecase
}

func NewCalendarHandler(calendarUsecase usecase.CalendarUsecase) CalendarHandler {
	return &calendarHandler{calendarUsecase: calendarUsecase}
}

type requestCalendar struct{}

type responseCalendar struct {
	Day   string `json:"day"`
	Plans []struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"plans"`
}

func (h *calendarHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("uuid")
		if uuid != config.GetUser(c).UUID {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid uuid")
		}

		year := c.Param("year")
		month := c.Param("month")
		user := config.GetUser(c)
		getCalendar, err := h.calendarUsecase.Get(
			year,
			month,
			&domain.User{
				ID:       user.ID,
				Password: user.Password,
			},
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseCalendar{}
		for _, calendar := range getCalendar {
			res = append(res, responseCalendar{
				Day: calendar.Day,
				Plans: []struct {
					Title string `json:"title"`
					Link  string `json:"link"`
				}{},
			})
			for _, plan := range calendar.Plans {
				res[len(res)-1].Plans = append(res[len(res)-1].Plans, struct {
					Title string `json:"title"`
					Link  string `json:"link"`
				}{
					Title: plan.Title,
					Link:  plan.Link,
				})
			}
		}

		return c.JSON(http.StatusOK, res)
	}
}
