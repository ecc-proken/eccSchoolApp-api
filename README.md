# SchoolApp-Api

某学校のデータを持ってくる API

## フレームワーク

https://echo.labstack.com/

## パッケージ

- https://pkg.go.dev/github.com/gocolly/colly
  - スクレイピング
- https://pkg.go.dev/github.com/joho/godotenv
  - .env 読み込み
- https://pkg.go.dev/github.com/golang-jwt/jwt
  - 認証

## 依存関係

```
+----------------+
|   interface    |
+-------+--------+
        |
+-------v--------+
|    usecase     |
+-------+--------+
        |
+-------v--------+
|     domain     |
+-------^--------+
        |
+-------+--------+
| infrastructure |
+----------------+
```
