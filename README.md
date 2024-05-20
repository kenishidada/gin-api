# golang ginのAPI作成

## 本リポジトリの目的

golang(gin)のAPI作成方法を一通り学習すること。

## やったこと

* Golang(Gin)のAPI作成方法を学習
* 3層アーキテクチャでの実装を学習
* gormを学習
* DBとの接続方法を学習
* テスト方法を学習
* .envを使用した設定ファイルの読み込みを学習
* ログイン機能を実装
* routingの方法、ミドルウェアの追加を学習

### todo

* 関数のテスト方法を学習
* go docの書き方を学習
* エラー後のエラー文がデフォルトのままのため、状況に応じて適切なものを返すように修正

## サーバー起動

* DBを起動

```shell
docker-compose up -d
```

* サーバー起動

```shell
air
```

## API操作方法

* POST /auth/register ユーザ登録

```curl
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{
    "email": "sample@sample.com",
    "password": "password"
}'
```

* POST /auth/login ユーザログイン

```curl
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{
    "email": "sample@sample.com",
    "password": "password"
}'
```

jwt tokenを取得

```go
return
"token": "xxx" 
```

* GET /items 全取得

```curl
curl -X GET http://localhost:8080/items
```

* POST /items 新規作成

xxxにtokenを入れる

```curl
curl -X POST http://localhost:8080/items \
-H "Content-Type: application/json" \
"Authorization: Bearer xxx" \
-d '{
    "name": "test_item1",
    "price": 2000,
    "description": "this is test_item1"
}'
```

GET /items/:id 詳細取得

```curl
curl -X GET http://localhost:8080/items/1 \
-H "Authorization: Bearer xxx" 
```

PUT /items/:id 更新

```curl
curl -X PUT http://localhost:8080/items/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer xxx" \
-d '{
    "name": "test_update_item1",
    "price": 10000,
    "description": "this is test_update_item1",
    "sold_out": true
}'
```

DELETE /items/:id 削除

```curl
curl -X DELETE http://localhost:8080/items/1
```

## テスト実行

* 以下コマンドを実行

```shell
go test ./...
```
