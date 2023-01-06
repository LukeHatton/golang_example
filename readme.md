[TOC]

# 随笔

本文档不记录成系统的笔记，只是学习golang过程中的一些随笔

## 查看go文档

### 使用`go doc`命令行查看

可以以包、类型、函数级别来分别查看文档

```shell
# 查看reflect包的文档
go doc reflect
# 查看reflect包中的函数的文档
go doc reflect.ValueOf
# 查看reflect.Value类型的文档
go doc reflect.Value
# 查看reflect.Value类型的Int方法的文档
go doc reflect.Value.Int
```

### 使用`godoc`在线浏览

通过godoc工具可以启动一个本地go文档服务器，方便在浏览器中直接浏览文档，这个文档要比官方API文档更全（因为包含用户本地库）

在go 1.13之前是自带godoc工具的，但是在之后的发行包中移除了，需要手动安装

```shell
# 可能需要先配置GOPROXY环境变量为国内代理，原始地址为https://goproxy.io
go env -w GOPROXY=https://goproxy.cn
go install golang.org/x/tools/cmd/godoc@latest
# 安装成功后，就可以启动并浏览本地在线go doc文档了
godoc -http=:8080
```

很多网上资料说安装godoc应该使用~~go get golang.org/x/tools/cmd/godoc~~，但是这个方法已过期，不适用于新版本的golang