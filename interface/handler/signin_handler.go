package handler

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
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
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (h *signinHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*config.JwtCustomClaims)
		id := claims.User.Id
		password := claims.User.Passwd

		fmt.Print(id, password)

		getSignin, err := h.signinUsecase.Get(&domain.User{
			Id:     id,
			Passwd: password,
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
