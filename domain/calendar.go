package domain

type Calendar struct {
	Day   string
	Plans Plans
}

type Plans struct {
	Title []string
	Link  []string
}
