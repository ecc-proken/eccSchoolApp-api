package infrastructure

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/domain/repository"
)

type SigninRepository struct{}

func NewSigninRepository() repository.SigninRepository {
	return &SigninRepository{}
}

func (r *SigninRepository) Get(user *domain.User) (*domain.Signin, error) {
	c := config.ECCLogin(user)

	// 返す値の初期化
	var title string

	// title取得
	c.OnHTML(".home_back", func(e *colly.HTMLElement) {
		title = e.Text
		fmt.Print(title)
	})

	// ログインページ
	c.Visit(os.Getenv("APP_DOMAIN") + os.Getenv("APP_LOGIN"))

	// titleに値が入っていなかったらログイン失敗
	if title == "" {
		return &domain.Signin{
			Status:  401,
			Message: "unauthorized error",
		}, nil
	}

	return &domain.Signin{
		Status:  200,
		Message: "success",
	}, nil
}
