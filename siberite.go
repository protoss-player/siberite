package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	service "github.com/bogdanovich/siberite/service"
)

var (
	dataDir     = flag.String("data", "./data", "path to data directory")
	hostAndPort = flag.String("listen", "0.0.0.0:22133", "ip and port to listen")
	versionFlag = flag.Bool("version", false, "prints current version")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	service := service.New(*dataDir)

	if *versionFlag {
		fmt.Println(service.Version())
		os.Exit(0)
	}

	laddr, err := net.ResolveTCPAddr("tcp", *hostAndPort)
	if nil != err {
		log.Fatalln(err)
	}

	go service.Serve(laddr)

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	// Stop the service gracefully.
	service.Stop()
}
