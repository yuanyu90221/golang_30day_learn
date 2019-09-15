package main

import (
	"fmt"
	"util/add"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(add.Add(1, 2))
	log.Info("雷大範例很好用")
}
