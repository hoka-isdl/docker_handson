#ベースイメージの取得
FROM golang:1.18.3-alpine
#パッケージの更新
RUN apk update
#作業ディレクトリの指定
WORKDIR /server

CMD ["go", "run", "main.go"]