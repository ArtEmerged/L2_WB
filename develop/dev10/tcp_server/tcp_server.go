package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

const listenMsg = "Listening on the port:"

func main() {
	port := flag.String("port", "5050", "--port 5050")
	flag.Parse()

	lis, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer lis.Close()
	fmt.Println(listenMsg + *port)
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		go Handler(conn)
	}
}

func Handler(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(os.Stdin, "joined the server:%s\n", conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Fprintf(os.Stdin, "left the server:%s\n", conn.RemoteAddr().String())
			} else {
				fmt.Fprint(os.Stderr, err)
			}
			return
		}
		fmt.Fprintf(os.Stdout, "client: %s", string(message))
		conn.Write([]byte("server: " + message))
	}
}
