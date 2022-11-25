package domain

// 出席情報
type Attendance struct {
	Title    string // タイトル
	Rate     string // 出席率
	Count    string // 出席数
	Absence  string // 欠席
	Lateness string // 遅刻
}
