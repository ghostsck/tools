package pool2;

import (
	"net"
	"io"
	"log"
)

func handler(conn net.Conn) {
	for {
		io.Copy(conn, conn);
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8080");
	if err != nil {
		log.Fatal(err);
	}

	for {
		conn, err := lis.Accept();
		if err != nil {
			continue;
		}
		go handler(conn);
	}
}