package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type TimetableHandler interface {
	Get() echo.HandlerFunc
}

type timetableHandler struct {
	timetableUsecase usecase.TimetableUsecase
}

func NewTimetableHandler(timetableUsecase usecase.TimetableUsecase) TimetableHandler {
	return &timetableHandler{timetableUsecase: timetableUsecase}
}

type requestTimetable struct{}

type responseTimetable struct {
	Date      string            `json:"date"`
	Weekday   string            `json:"weekday"`
	Timetable []timetableDetail `json:"timetable"`
}

type timetableDetail struct {
	Period       string `json:"period"`
	SubjectTitle string `json:"subjectTitle"`
	Classroom    string `json:"classroom"`
	Teacher      string `json:"teacher"`
}

func (h *timetableHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("uuid")
		if uuid != config.GetUser(c).UUID {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid uuid")
		}

		week := c.Param("week")
		weekInt, err := strconv.Atoi(week)
		user := config.GetUser(c)
		getTimetable, err := h.timetableUsecase.Get(
			weekInt,
			&domain.User{
				ID:       user.ID,
				Password: user.Password,
			},
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTimetable{
			Date:      getTimetable.Date,
			Weekday:   getTimetable.Weekday,
			Timetable: []timetableDetail{},
		}
		for _, v := range getTimetable.Timetable {
			res.Timetable = append(res.Timetable, timetableDetail{
				Period:       v.Period,
				SubjectTitle: v.SubjectTitle,
				Classroom:    v.Classroom,
				Teacher:      v.Teacher,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}
