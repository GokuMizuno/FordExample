//Gordon Stangler
//Get all IP addresses a machine may have

package main

import(
//	"os"	//os not used, per GO compiler
	"net"
	"fmt"
//	"log"	//log not used, per GO compiler
)

func main(){
	ifaces, err := net.Interfaces()
//	if err != nil {  log.Fatal(err)  }
	if err != nil {  fmt.Printf("Error: %s", err)  }

	for _, i := range ifaces {
		address, err := i.Addrs()
		for _, addr := range address{
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
//			fmt.Printf("IP address %v\n", ip)
		}
	}
}
