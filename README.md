# abobon-jr

abobon-jrはDiscord Botの名前です。このリポジトリにはabobon-jrのソースコードが含まれています。

## 開始方法

### 1. Docker Desktopをインストールする

まず、Docker Desktopをインストールしてください。インストール方法は[公式サイト](https://docs.docker.com/desktop/install/mac-install/)を参照してください。

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


# Branch戦略

## 1. Issue管理

- GItHub Projectを用いて管理する
- 各Issueには、タスクの詳細な説明や完了条件が記載されている


## 2. ブランチ運用ルール

本プロジェクトでは以下のルールにしたがってブランチ運用を行う
### ブランチの種類

- **main**: 本番環境のブランチ
    - ここには直接コミットしない
    - mainブランチでは開発は行わない

- **feature**: 機能の追加や変更、不具合の修正を実際に開発を行うブランチ
    - 作業完了後、mainにマージする
- **hotfix**: 本番環境に反映後にバグが発生した場合、修正を行うブランチ
    - mainブランチから作成し、作業完了後、mainブランチにマージする
 

### 基本的な開発の流れ

1. mainブランチからfeatureブランチを作成する
2. featureブランチで開発をする
3. 2が完了したら、mainブランチにマージする
4. デプロイをする


### バグが発生した場合

1. mainブランチからhotfixブランチを作成する
2. hotfixブランチで開発をする
3. 2が完了したら、mainブランチにマージする
4. デプロイをする


### ブランチの作成

- Issueから新しいブランチを作成する
- ブランチ名は以下の命名規則に従う：
    - [ブランチの種類]/[機能の特徴]/[Issue番号]
    - 例
        - feature/slash-command/#xx
        - hotfix/fix-get-thesis/#xx


## 3. 開発とコミット

- 作成したブランチで開発を行う
- コミットメッセージは何を変更したのかを明確に記述してください


## 4. Pull Requestの作成

- 開発が完了したら、mainブランチに対してPRを作成する
- PRタイトルには、関連するIssue番号を含める：
    - [Issue番号][簡潔な変更内容]
    - 例：#xx fix user authentication


## 5. コードレビューとマージ

- 可能なかぎりレビューを受ける
-  レビューを通過したら、レビュワーがPRをマージする


## 6. タスクの完了

- PRがマージされたら、対応するIssueをクローズする
- GitHub ProjectsのstatusをDoneに変更する
