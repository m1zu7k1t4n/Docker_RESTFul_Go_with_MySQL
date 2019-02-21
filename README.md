# Docker on RESTFul Go with MySQL

Docker上でRESTFulなGoサーバーを立ち上げるハンズオン

## 遊び方

こちらのリポジトリの遊び方をQiitaに記事にしてあります

`https://qiita.com/m1zu7k1t4n/items/f3625e96cf43abaac55f`

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
