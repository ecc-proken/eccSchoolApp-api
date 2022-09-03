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

## systemd

```sh
su -

cd /root/go/src/eccSchoolApp-api
cp .env.example .env
make build

cat << EOF > /lib/systemd/system/eccSchoolApp-api.service
[Unit]
Description = eccSchoolApp-api.service daemon

[Service]
Environment="GOPATH=/root/go"
WorkingDirectory=/root/go/src/eccSchoolApp-api
ExecStart=/root/go/src/eccSchoolApp-api/build/eccSchoolApp-api
Restart=always
Type=simple
User=root

[Install]
WantedBy = multi-user.target
EOF

systemctl daemon-reload
systemctl restart eccSchoolApp-api.service
systemctl enable eccSchoolApp-api.service
```

## varnish

```sh
sub vcl_recv {
  if (req.url ~ "/signin" || req.url ~ "/uuid") {
    return (pipe);
  }
  return (hash);
}

sub vcl_backend_response {
  if (! (beresp.status == 200)) {
    set beresp.ttl = 60s;
  }
}
```