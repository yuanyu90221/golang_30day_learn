# day2_log

## focus

  day2 的重點在於

  1 如何 建立自己的lib

    並且在自己的 go func 內使用

  2 如何 使用別人的libary

### 1 build own library

  要讓自己的library可以 被import

  達成的條件 就是要在 $GOPATH/src 下面

  建立自己的library

  或是建置完後 把自己的 library 透過 install方式

  把自己的lib 建置到$GOPATH/bin

### 2 import from reference dir

```golang==
  import (
    "package相對於$GOPATH/src的路徑"
  )
```

### 3 install other lib

```golang==
go get "lib的url"
```