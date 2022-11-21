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

## example

```shell
# login
curl -X POST -H "Content-Type: application/json" -d '{"id":"2200000", "pw":"hoge"}' https://eccschoolapp-api.yumekiti.net/signin

# uuid
curl https://eccschoolapp-api.yumekiti.net/uuid -H "Authorization: Bearer <token>"

# news
curl https://eccschoolapp-api.yumekiti.net/<uuid>/news -H "Authorization: Bearer <token>"
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
    set beresp.uncacheable = true;
    return (deliver);
  }
  set beresp.ttl = 3600s;
}

sub vcl_deliver {
  set resp.http.Access-Control-Allow-Origin = "*";

  if (req.method == "OPTIONS") {
    set resp.http.Access-Control-Allow-Methods = "GET, POST, OPTIONS";
    set resp.http.Access-Control-Allow-Headers = "*";

    set resp.status = 204;
  }
}
```

## テスト

```sh
make test
```

### モック
  
```sh
docker run -v "$PWD":/src -w /src vektra/mockery --all --keeptree
sudo chown -R $USER:$USER ./mocks
```
