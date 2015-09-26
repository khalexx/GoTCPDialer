package main

import (
	"bufio"
	"log"
	"os"
)

type SimpleUserConsole struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (s *SimpleUserConsole) Init() {
	s.Reader = bufio.NewReader(os.Stdin)
	s.Writer = bufio.NewWriter(os.Stdout)
}

func (s *SimpleUserConsole) Read() (str string) {
	str, err := s.Reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return str
}

func (s *SimpleUserConsole) Write(str string) {
	_, err := s.Writer.WriteString(str)
	defer s.Writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
