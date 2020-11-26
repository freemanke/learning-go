# Learning golang

## 包管理
v1.12 版本后，golang 引入了 go modules 来作为官方包管理
```
go mod init freemanke/learning-go/v2
go get -u github.com/gin-gonic/gin

```

## go get 代理设置
go get 如果没有代理，会非常慢，无法拉取任何包，代理设置方法：在环境变量新增http_proxy, 值为http://127.0.0.1:1080

## 开发框架
- web framework - gin


## TDD
