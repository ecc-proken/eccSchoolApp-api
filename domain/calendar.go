package domain

// カレンダー
type Calendar struct {
	Day   string  // 日付
	Plans []Plans // 予定
}

type Plans struct {
	Title string
	Link  string
}
