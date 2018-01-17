package main

import (
	"fmt"
	"os"
    "github.com/softinno/goutils/sockt"
)


func main() {
	
	testGet()
	// testTunnel()
}

func testTunnel() {
	conn, _ := sockt.CreateConn("localhost:8443")
	defer conn.Close()
	sockt.DoTunnel(conn, "127.0.0.1", 9121, "", "")
}

func testGet() {
	conn, err := sockt.CreateConn("localhost:8443")
	defer conn.Close()
	if err != nil {
		fmt.Println("Error on connection creation: " + err.Error())
		os.Exit(1)
	}
	
	conn, err = sockt.DoTunnel(conn, "127.0.0.1", 80, "", "")
	if err != nil {
		fmt.Println("Error on tunnel creation: " + err.Error())
		os.Exit(1)
	}
	
	result, err := sockt.Get(conn, "/");
	if err != nil {
		fmt.Println("Error on get: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Get result:")
	fmt.Println(string(result))
}
