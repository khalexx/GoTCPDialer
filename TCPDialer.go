package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	var (
		console SimpleUserConsole
		sent    string
	)

	console.Init()

	console.Write("Connecting\n")
	conn, err := net.Dial("tcp4", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	connWriter := bufio.NewWriter(conn)
	connReader := bufio.NewReader(conn)

	sent = console.Read()

	console.Write("Sending " + sent + "\n")

	connWriter.WriteString(sent)
	connWriter.Flush()

	result, err := connReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	console.Write("close connection and print result\n")
	console.Write(result + "\n")
	conn.Close()
}
