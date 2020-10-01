package main

import (
	"fmt"
	"github.com/sivaramsajeev/go-starter/dirs"
	"github.com/sivaramsajeev/go-starter/internal/auth"
	"gopkg.in/alecthomas/kingpin.v2"
)


var (
  debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
  timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
  ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
  count   = kingpin.Arg("count", "Number of packets to send").Int()
)

func main() {
	fmt.Println("Gopher z Just starting out :)")
	fmt.Println(auth.GetAuth())
	fmt.Println(dirs.AddonPrint())


  kingpin.Version("0.0.1")
  kingpin.Parse()
  fmt.Printf("Would ping: %s with timeout %s and count %d\n", *ip, *timeout, *count)
}