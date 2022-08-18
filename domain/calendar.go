package domain

type Calendar struct {
	Day   string
	Plans struct {
		Title []string
		Link  []string
	}
}
