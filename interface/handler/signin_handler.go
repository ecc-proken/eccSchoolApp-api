package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type SigninHandler interface {
	Get() echo.HandlerFunc
	Mock() echo.HandlerFunc
}

type signinHandler struct {
	signinUsecase usecase.SigninUsecase
}

func NewSigninHandler(signinUsecase usecase.SigninUsecase) SigninHandler {
	return &signinHandler{signinUsecase: signinUsecase}
}

type requestSignin struct{}

type responseSignin struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (h *signinHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("uuid")
		if uuid != config.GetUser(c).UUID {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid uuid")
		}

		user := config.GetUser(c)
		getSignin, err := h.signinUsecase.Get(&domain.User{
			ID:       user.ID,
			Password: user.Password,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseSignin{
			Status:  getSignin.Status,
			Message: getSignin.Message,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (h *signinHandler) Mock() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := responseSignin{
			Status:  200,
			Message: "いらっしゃいませご主人様^~",
		}

		return c.JSON(http.StatusOK, res)
	}
}