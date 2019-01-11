package main

import (
	"./business_server"
	"./caddy_server"
)

func main() {
	caddyInst := caddy_server.GetServer()
	business_server.StartHTTPServer()
	caddyInst.Wait()
}
