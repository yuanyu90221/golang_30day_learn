# goModule
[ref doc](https://ithelp.ithome.com.tw/articles/10217414)
## introduction

為了解決 專案goPath問題

go modules有build cache, 在CI build server上速度飛快.

所以出現了 go modules 解決 dependency問題


## prerequest

go >= 1.11

## how to use
1 create your module directory

2 use go mod init with your module directory
```golang=
go mod init ${folder}/${moduleName}
```
## setup go.mod

```mod=
module day12_gomodule/goModule

go 1.12

require github.com/sirupsen/logrus v1.4.2
```

## use go mod download to download dependency your need

```golang=
go mod download
```