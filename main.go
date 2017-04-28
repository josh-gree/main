package main

import (
	"github.com/josh-gree/serverbase"
	"gopkg.in/alecthomas/kingpin.v2"
)

var master = kingpin.Arg("master", "master process?").Bool()

func main(){

	kingpin.Parse()
	serverbase.Listen(*master)
}
