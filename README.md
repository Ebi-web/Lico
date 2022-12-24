# Lico

# 環境構築

## ホットリロード

Airを使用。

## デバッグ

1. ngrok 8080をリッスンしてそのURIを[LINE Developers Console](https://developers.line.biz/console/channel/1657725137/messaging-api)にてWebhook URLに登録。
2. ただし、ホットリロードを効かせないと `go run index.go`した時のコンパイル結果が利用されてしまうので、必ず `air`でホットリロードすること。 `air` しておくと勝手にmainパッケージのmain関数が実行されてサーバが立つので便利。
3. あとはLINEから好きにメッセージを送ればローカルで確認できる。