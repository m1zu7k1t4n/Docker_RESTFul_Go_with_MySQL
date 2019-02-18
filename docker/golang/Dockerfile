FROM golang:1.11-alpine
LABEL maintainer "arasan01 <sinzan062871@gmail.com>"
RUN set -x && \
  apk add git; \
  go get -u github.com/gin-gonic/gin; \
  go get -u github.com/gin-contrib/cors; \
  go get -u github.com/jinzhu/gorm; \
  go get github.com/go-sql-driver/mysql

