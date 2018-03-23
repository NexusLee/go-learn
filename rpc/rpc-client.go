package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Response struct {
	Country   string
	Province  string
	City      string
	ISP       string
	Latitude  float64
	Longitude float64
	TimeZone  string
}
type Agrs struct {
	IpString string
}

func main() {
	client, err := jsonrpc.Dial("tcp", "192.168.0.101:8001")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	var res Response
	err = client.Call("Ip2addr.Address", Agrs{"your ip address"}, &res)
	if err != nil {
		log.Fatal("Ip2addr error:", err)
	}
	fmt.Println(res)

}