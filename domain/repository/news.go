package repository

import "github.com/yumekiti/eccSchoolApp-api/domain"

type NewsRepository interface {
	Get() (*domain.News, error)
}
