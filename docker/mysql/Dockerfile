# 基本となるMySQLのイメージを指定
FROM mysql:8.0.29

# カスタム設定ファイルをイメージにコピー
COPY mysql/my.cnf /etc/mysql/my.cnf

# 初期スクリプトフォルダにSQLファイルをコピー
COPY mysql/migrations/ /docker-entrypoint-initdb.d/

# ヘルスチェックの設定
HEALTHCHECK --interval=5s --timeout=5s --retries=5 --start-period=5s \
    CMD mysqladmin ping -h 127.0.0.1 -u root || exit 1
