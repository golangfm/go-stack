package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mholt/caddy"
	_ "github.com/mholt/caddy/caddyhttp"
)

func init() {
	caddy.SetDefaultCaddyfileLoader("default", caddy.LoaderFunc(loadConfig))
}

func loadConfig(serverType string) (caddy.Input, error) {
	contents, err := ioutil.ReadFile(caddy.DefaultConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	fmt.Printf("defaultConfiFile:%s\n", caddy.DefaultConfigFile)
	fmt.Printf("loading Caddyfile: %s\n", string(contents))
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}

func main() {
	caddy.AppName = "MyApp"
	caddy.AppVersion = "0.1"
	caddyfile, err := caddy.LoadCaddyfile("http")
	if err != nil {
		panic(err)
	}
	inst, err := caddy.Start(caddyfile)
	if err != nil {
		panic(err)
	}
	inst.Wait()
}
