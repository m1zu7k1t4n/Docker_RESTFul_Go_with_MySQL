# Docker on MySQL and Some Connector Client

MySQLとその他の言語を組み合わせてデータベースを用いた環境を構築する

# Commands

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
docker exec -it python_host sh
docker exec -it golang_host sh
```
