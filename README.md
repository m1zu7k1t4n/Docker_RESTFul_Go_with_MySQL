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

## reference

Build RESTful API service in golang using gin-gonic framework
`https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3`
