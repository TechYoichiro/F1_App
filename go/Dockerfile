# ベースイメージを指定
FROM golang:latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /api/src

# ローカルのソースコードをコンテナにコピー
COPY . .

# 必要なパッケージをインストール
RUN go mod download

# アプリケーションをビルド
RUN go mod tidy

# 公開予定のコンテナのポートを明示
EXPOSE 8080

# アプリケーションを実行
CMD ["go", "run", "cmd/main.go"]
