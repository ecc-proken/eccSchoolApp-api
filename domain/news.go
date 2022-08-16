package domain

type News struct {
	NewsDetail []NewsDetail
}

type NewsDetail struct {
	ID    int
	Title string
	Date  string
	Link  string
}
