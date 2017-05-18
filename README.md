## 项目简介
ElasticHD 是ElasticSearch的可视化DashBoard， 支持Es状态监控、实时搜索、Index Template 快捷替换&查看、支持Index list、支持sql 查询。

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


