# shanyan-go

```
go get -u github.com/ningge123/shanyan-go
```

## use

### 一键登录V2（获取手机号码）接口

```go

package main

import (
	"fmt"
	"log"
	"github.com/ningge123/shanyan-go"
)

func main()  {
	client := shanyan.NewClient("app ID", "app key")
	mobile, err := client.MobileQuery("token")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(mobile)
}


```