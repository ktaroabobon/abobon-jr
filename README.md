# abobon-jr

abobon-jrはDiscord Botの名前です。このリポジトリにはabobon-jrのソースコードが含まれています。

## 開始方法

### 1. Dockerをインストールする

まず、Dockerをインストールしてください。インストール方法は[公式サイト](https://www.docker.com/get-started)を参照してください。

### 2. リポジトリのクローン

次に、このリポジトリをクローンします。

```sh
git clone https://github.com/ktaroabobon/abobon-jr.git

cd abobon-jr
```


### 3. 初期設定

初期設定を行うために、以下のコマンドを実行します。
```sh
make init
```

### 4. Botの実行

Botを実行するには、以下のコマンドを実行します。

```sh
make run
```

## ユーティリティ
### ログの確認
#### アプリケーションのログ

```sh
make app/logs
```

##### データベースのログ

```sh
make db/logs
```

### コンテナへのアクセス

#### アプリケーションコンテナ

```sh
make app/login
```

#### データベースクライアント

```sh
make db/client
```

## 補足

以上がabobon-jrの基本的な開発方法です。詳細な設定やカスタマイズについては、ソースコードを参照してください。