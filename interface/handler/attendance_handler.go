package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type AttendanceHandler interface {
	Get() echo.HandlerFunc
}

type attendanceHandler struct {
	attendanceUsecase usecase.AttendanceUsecase
}

func NewAttendanceHandler(attendanceUsecase usecase.AttendanceUsecase) AttendanceHandler {
	return &attendanceHandler{attendanceUsecase: attendanceUsecase}
}

type requestAttendance struct{}

type responseAttendance struct {
	Title    string `json:"title"`
	Rate     string `json:"rate"`
	Count    string `json:"count"`
	Absence  string `json:"absence"`
	Lateness string `json:"lateness"`
}

func (h *attendanceHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("uuid")
		if uuid != config.GetUser(c).UUID {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid uuid")
		}

		user := config.GetUser(c)
		getAttendance, err := h.attendanceUsecase.Get(
			&domain.User{
				ID:       user.ID,
				Password: user.Password,
			},
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseAttendance{}
		for _, attendance := range getAttendance {
			res = append(res, responseAttendance{
				Title:    attendance.Title,
				Rate:     attendance.Rate,
				Count:    attendance.Count,
				Absence:  attendance.Absence,
				Lateness: attendance.Lateness,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}
