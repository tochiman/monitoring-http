# ステージ1(build)
# Goのイメージを指定
FROM golang:1.19-alpine as go
ENV CGO_ENABLED 0
# ローカルにあるコードをコンテナ内に格納する場所に作成と移動
WORKDIR /go/src
# コンテナ内のアップデートとGitを使用するためダウンロード
RUN apk update && apk add git
# ローカルにあるファイルと
COPY ./ ./
RUN go mod download && go build -o /go/bin/main ./main.go

# ステージ2(Goの実行環境)
# 軽量なディストローレスのDebianを指定
FROM gcr.io/distroless/static-debian11
# go(ステージ1)からビルドされた「mainファイル」をステージ1からステージ2の「/」にコピー
COPY --from=go /go/bin/main /
# metrics情報を8080番ポートで公開されているため
EXPOSE 8080
# 監視ソフトの起動
CMD ["/main"]