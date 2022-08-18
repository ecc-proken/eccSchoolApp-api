package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type SigninHandler interface {
	Get() echo.HandlerFunc
}

type signinHandler struct {
	signinUsecase usecase.SigninUsecase
}

func NewSigninHandler(signinUsecase usecase.SigninUsecase) SigninHandler {
	return &signinHandler{signinUsecase: signinUsecase}
}

type requestSignin struct{}

type responseSignin struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (h *signinHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		getSignin, err := h.signinUsecase.Get(&domain.User{
			Id:     os.Getenv("TEST_ID"),
			Passwd: os.Getenv("TEST_PW"),
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
