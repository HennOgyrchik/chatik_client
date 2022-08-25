package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	//подключение к  серверу
	conn, err := net.Dial("tcp", "192.168.0.123:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go writer(conn, &wg) //send to server

	wg.Add(1)
	go servAnsw(conn, &wg) //read of answer

	wg.Wait()
}

func servAnsw(conn net.Conn, wg *sync.WaitGroup) {
	for {
		buff := make([]byte, 1024)

		n, err := conn.Read(buff)
		if err != nil {
			break
		}

		fmt.Print(string(buff[0:n]))
	}
	defer wg.Done()
}

func writer(conn net.Conn, wg *sync.WaitGroup) {

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		if text == "" {
			text = " "
		}
		conn.Write([]byte(text))

	}
	defer wg.Done()
}
