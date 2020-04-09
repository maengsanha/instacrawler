// Package main runs Instagram crawler api server.
package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
