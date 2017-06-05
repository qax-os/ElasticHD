ElasticHD
-----------
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/farmerx/elasticHD/main)
> ElasticHD 是一款 ElasticSearch的可视化应用。不依赖ES的插件安装，更便捷；导航栏直接填写对应的ES IP和端口就可以操作Es了。目前支持如下功能：
 * ES 实时搜索
 * ES DashBoard 数据可视化
 * ES Index Template (在线修改、查看、上传）
 * SQL Converts to DSL
 * ES 基本查询文档
 
## ElasticHD应用页面
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120044177.png)
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120114338.png)
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120147401.png)

## 前端页面查看
```
go get https://github.com/farmerx/ElasticHD
# install dependencies
npm install

# serve with hot reload at localhost:8090
npm run dev
```

## ElasticHD build set
```
go get https://github.com/farmerx/ElasticHD
# 进入程序源码目录
cd ElasticHD
# install dependencies
npm install
# build vue 源码
npm run build
# 进入到服务端程序目录
cd main
# 使用statik 压缩编译好的程序
statik -src=../dist
# go build
GO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o elasticHD.exe github.com/elasticHD/main
```

## 二进制文件下载

选择所需要的二进制版本进行下载
* 请下载[download windows 64位](https://github.com/farmerx/ElasticHD/raw/master/elasticHD_windows_64.exe)
* 请下载[download mac_OSX 64位](https://github.com/farmerx/ElasticHD/raw/master/elasticHD_OSX)
* 请下载[download linux 64位](https://github.com/farmerx/ElasticHD/raw/master/elasticHD_linux)

## 二进制版本运行
> 命令行运行二进制软件， 然后打开浏览器访问127.0.0.1:9800



