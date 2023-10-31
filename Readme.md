# internship-go-api base

## 概要

DDD クリーンアーキテクチャでプロジェクトを学習するためのベースリポジトリです。  
内容やドキュメントに関してはフィードバックを受けて随時更新していきます。

## ディレクトリ構成

### /application

ビジネスロジックを配置します。

- /application/user
- /application/admin

のようにアクターごとにディレクトリを作成していきます。

### /cmd

アプリケーションのエントリーポイントを配置します。

### /domain

ドメイン層のコードを配置します。

### /infrastructure

インフラストラクチャ層のコードを配置します。  
各種clientやrepositoryはdomain層で定義されているinterfaceに依存させる形で実装します。

### /di

DIコンテナの設定ファイルを配置します。

### /config

dbへの接続情報等の設定ファイルを配置します。

### /migrations

migrationファイルを配置します。

## 環境構築

### config/config.example.yamlをコピー

```sh
cp config/config.example.yaml config/config.yaml
```

### .envrc.exampleをコピー

direnvを利用しています。

```sh
cp .envrc.example .envrc
direnv allow
```

### CLIをインストール

```sh
# DIコンテナを生成するためのツール
go install github.com/google/wire/cmd/wire@latest
# ホットリロードをするためのツール
go install github.com/cespare/reflex@latest 
# デバッガ
go install github.com/go-delve/delve/cmd/dlv@latest 
# 静的解析ツール
go install honnef.co/go/tools/cmd/staticcheck@latest 
# migrationツール
go install github.com/pressly/goose/v3/cmd/goose@latest 
```

### Migrationについて

gooseを利用してます Makefileにコマンドを切っています

```sh
# migrationの状態を確認する
make migrate-status

# migrationを反映する
make migrate-up

# migrationを1つrollbackする
make migrate-down

# migrationすべてrollbackする
make migrate-reset

# migrationをすべてrollbackして、再度migrationを反映する
make migrate-refresh

# sql形式のmigrationファイルを作成する
make migrate-create ARG=migration_file_name

# go形式のmigrationファイルを作成する
make migrate-create-go ARG=migration_file_name
```