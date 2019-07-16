// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type filesFlag []string

func (i *filesFlag) String() string {
	return "service file(s) containing struct"
}

func (i *filesFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var options struct {
	service      string
	serviceFiles filesFlag
	structName   string
	iface        string
	icomment     string
	pkgName      string
	outputFile   string
	comment      string
}

// Prints an error message and exits.
func fatal(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "error: %s\n", msg)
	os.Exit(1)
}

// Set up flags
func init() {
	flag.StringVar(&options.service, "svc", "", "service name")
	flag.Var(&options.serviceFiles, "sf", "service file(s) containing struct")
	flag.StringVar(&options.structName, "s", "Service", "struct to generate interface from")
	flag.StringVar(&options.iface, "i", "Servicer", "name of generated interface")
	flag.StringVar(&options.icomment, "ic", "Servicer represents the interface for implementing all endpoints for this service", "name of generated interface")
	flag.StringVar(&options.pkgName, "p", "", "package name of generated interface")
	flag.StringVar(&options.outputFile, "o", "", "output file name. Print to stdout if not provided")
	flag.Parse()
}

func main() {
	var err error
	outFile := filepath.Join(options.service, "interface.go")
	args := []string{}
	for _, v := range options.serviceFiles {
		file := filepath.Join(options.service, v)
		fmt.Printf("generating from %s\n", file)
		args = append(args, "-f", file)
	}
	args = append(args, "-s", options.structName, "-i", options.iface, "-p", options.pkgName, "--iface-comment", options.icomment, "-c", "Code generated by gen_interface.go. DO NOT EDIT.", "-o", outFile)
	cmd := exec.Command("ifacemaker", args...)
	_, err = cmd.Output()
	if err != nil {
		fatal("%v", err)
	}
}
