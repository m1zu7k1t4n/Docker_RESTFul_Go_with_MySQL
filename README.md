# Docker on MySQL and Connector Client Application

MySQLとその他の言語を組み合わせてデータベースを用いた環境を構築する

## Commands

起動

```sh
docker-compose up -d
```

MySQLサーバーに接続

```sh
mysql -hlocalhost -utest -ptest test
```

それぞれのコンテナに接続

```sh
docker-compose exec {serviceName} sh
```
