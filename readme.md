# ![RealWorld Example App](logo.png)


[![Build Status](https://travis-ci.org/wangzitian0/golang-gin-starter-kit.svg?branch=master)](https://travis-ci.org/wangzitian0/golang-gin-starter-kit)
[![codecov](https://codecov.io/gh/wangzitian0/golang-gin-starter-kit/branch/master/graph/badge.svg)](https://codecov.io/gh/wangzitian0/golang-gin-starter-kit)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/wangzitian0/golang-gin-starter-kit/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/wangzitian0/golang-gin-starter-kit?status.svg)](https://godoc.org/github.com/wangzitian0/golang-gin-starter-kit)

> ### Golang/Gin 代码库包含工业生产的样例（CRUD, auth, advanced patterns, etc) 遵循 [RealWorld](https://github.com/make4yourself/realworld) 规范和 API。


创建这个代码库是为了掩饰一个使用 **Golang/Gin** 构建的完全成熟全栈应用程序，包含 CRUD 操作，认证，路由，分页等。


# 它是如何工作的
```
.
├── gorm.db
├── hello.go
├── common
│   ├── utils.go        //一些小的工具函数
│   └── database.go     //DB 链接管理
└── users
    ├── models.go       //数据模型的定义 & DB 操作
    ├── serializers.go  //响应处理 & 格式
    ├── routers.go      //业务逻辑 & 路由绑定
    ├── middlewares.go  //响应前处理 & 响应后置处理
    └── validators.go   //form/json 检查器
```

# 做好准备

## 安装 Golang
https://golang.google.cn/doc/install
## 环境配置
请确保配置好了 ~/.*shrc 下的变量
```
➜  echo $GOPATH
/Users/user/test/
➜  echo $GOROOT
/usr/local/go/
➜  echo $PATH
...:/usr/local/go/bin:/Users/user/test//bin:/usr/local/go//bin
```
## Install Go mod
go mod 是 1.11 版本之后官方自带的包管理工具

[go mod资料](https://www.jianshu.com/p/bbed916d16ea)
```
cd
go get 
```

## 开始
```
➜  go run hello.go
```

## 测试
```
➜  go test -v ./... -cover
```

## 代办
- 更优雅的配置
- 测试覆盖 (common & users 100%, article 0%)
- ProtoBuf 支持
- 代码结构优化（我觉得有些地方可以使用Interface替换）
- 持续集成（已完成）
- 输出教程
