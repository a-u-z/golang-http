package main

import "log"

var webPort = ":8123"

func main() {
	s := NewServer()

	err := s.router.Run(webPort)
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Panic(err)
	}
}
