ElasticHD
-----------
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/farmerx/elasticHD/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Luxurioust/aurora)](https://goreportcard.com/report/github.com/farmerx/elasticHD/main)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/Luxurioust/aurora/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/elasticHD/elasticHD.svg?label=Release)](https://github.com/farmerx/elasticHD/releases)
> ElasticHD 是一款 ElasticSearch的可视化应用。不依赖ES的插件安装，更便捷；导航栏直接填写对应的ES IP和端口就可以操作Es了。目前支持如下功能：
 * ES 实时搜索
 * ES DashBoard 数据可视化
 * ES Index Template (在线修改、查看、上传）
 * SQL Converts to DSL
 * ES 基本查询文档
 

## Installation

[Precompiled binaries](https://github.com/farmerx/elasticHD/releases) for supported operating systems are available.

## Contributing

Contributions are welcome! Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

## Get Start
```
# 需要go环境
go get https://github.com/farmerx/ElasticHD
# 进入到应用目录下
cd ElasticHD
npm install
# build vue 源码
npm run build
# 进入到服务端程序目录
cd ./main
# 使用statik 压缩编译好的程序(github上的一款go应用)
statik -src=../dist
# go build
GO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o elasticHD.exe github.com/elasticHD/main
```
## ElasticHD应用页面
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120044177.png)
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120114338.png)
![image](https://github.com/farmerx/ElasticHD/blob/master/snp20170518120147401.png)
![image](https://raw.githubusercontent.com/farmerx/ElasticHD/master/WX20170605-105148.png)

## Todo
...


## Licenses

This program is under the terms of the MIT License. See [LICENSE](https://github.com/farmerx/elasticHD/blob/master/LICENSE) for the full license text.


