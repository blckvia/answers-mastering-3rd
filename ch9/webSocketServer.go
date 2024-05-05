package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func echo(c net.Conn, numberToGenerate int) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	arr := make([]byte, numberToGenerate)
	for i := 0; i < numberToGenerate; i++ {
		random := rng.Intn(256)
		arr = append(arr, byte(random))
	}

	_, err := c.Write(arr)
	if err != nil {
		fmt.Println("Write:", err)
		return
	}
}

func main() {
	// Read socket path
	if len(os.Args) == 1 {
		fmt.Println("Need socket path and number to generate")
		return
	}

	socketPath := os.Args[1]
	numberToGenerate := os.Args[2]
	nGen, err := strconv.Atoi(numberToGenerate)
	if err != nil {
		fmt.Println("Not a valid number:", err)
		return
	}

	// If socketPath exists, delete it
	_, err = os.Stat(socketPath)
	if err == nil {
		fmt.Println("Deleting existing", socketPath)
		err := os.Remove(socketPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			return
		}
		go echo(fd, nGen)
	}
}
