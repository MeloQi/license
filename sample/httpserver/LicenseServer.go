package main

import (
	"flag"
	"github.com/MeloQi/license/httpapi"
	"log"
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ip := flag.String("i", "", "Server ip")
	port := flag.String("p", "8010", "Server port")
	flag.Parse()
	addr := *ip + ":" + *port

	HttpApi := httpapi.GetGenLicHttpApiInst(addr)
	HttpApi.Start()
	log.Println("LicenseServer Start OK, Listen On ", addr)
	for {
		time.Sleep(time.Minute)
	}
}
