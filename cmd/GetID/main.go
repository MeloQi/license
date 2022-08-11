package main

import (
	"flag"
	"fmt"
	"github.com/MeloQi/license/machineid"
)

var appid = ""

func main() {
	flag.StringVar(&appid, "appid", "", "应用id")
	flag.Parse()
	if appid == "" {
		return
	}
	if id, err := machineid.GetMachineid(appid); err == nil {
		fmt.Println(id)
	}
}
