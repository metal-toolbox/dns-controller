// Package main runs the dns controller
package main

//go:generate sqlboiler crdb --add-soft-deletes

import "go.hollow.sh/dnscontroller/cmd"

func main() {
	cmd.Execute()
}
