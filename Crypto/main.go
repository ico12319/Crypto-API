package main

import "CryptoToken/pkg/server"

func main() {
	serv := server.NewServer()
	serv.Start()
}
