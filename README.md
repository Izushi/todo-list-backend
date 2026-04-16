# todo-list-backend

Go + PostgreSQL で作る TODO リストの REST API バックエンド (学習用)。

## 必要なもの

- [mise](https://mise.jdx.dev/) (Go と golang-migrate のバージョン管理。mise.toml 通りに自動で入る)
- [Docker](https://www.docker.com/) + Docker Compose

## セットアップ

```sh
# 1. 環境変数ファイルを用意
cp .env.example .env

# 2. ツール (Go, gomigrate) を mise でインストール
mise install

# 3. 依存パッケージを取得
mise run tidy

# 4. PostgreSQL を起動
mise run db-up

# 5. マイグレーションを適用 (テーブル作成)
mise run migrate-up

# 6. サーバー起動 (DB接続確認)
mise run run
```

## ディレクトリ構成

```
.
├── cmd/server/        # エントリーポイント (main.go)
├── internal/          # アプリ内部のコード (これから書く)
│   ├── handler/       # HTTPハンドラ (リクエスト受付)
│   ├── repository/    # DBアクセス層
│   └── model/         # 構造体 (User, Task, Category)
├── migrations/        # SQLマイグレーション
├── docker-compose.yml # PostgreSQL定義
└── mise.toml          # ツールバージョン + タスク定義
```

## よく使うコマンド

```sh
mise tasks                            # タスク一覧
mise run db-up                        # DB起動
mise run db-shell                     # psqlでDBに入る
mise run migrate-up                   # マイグレーション適用
mise run migrate-down                 # 直近1つをロールバック
mise run migrate-create add_something # 新しいマイグレーションファイル作成
mise run run                          # サーバー起動
```

`mise run` は `mise r` と省略可能。

## API エンドポイント (計画)

### Users
- `POST /users` - ユーザー作成
- `GET /users/{userId}` - ユーザー取得

### Categories
- `POST /categories` - カテゴリ作成
- `GET /categories` - カテゴリ一覧
- `DELETE /categories/{categoryId}` - カテゴリ削除

### Tasks
- `POST /tasks` - タスク作成
- `GET /tasks?taskStatus=todo&categoryId=1` - タスク一覧 (フィルタ対応)
- `GET /tasks/{taskId}` - タスク詳細
- `PUT /tasks/{taskId}` - タスク全体更新
- `PATCH /tasks/{taskId}` - タスク部分更新
- `DELETE /tasks/{taskId}` - タスク削除
