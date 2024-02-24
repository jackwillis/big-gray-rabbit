package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	domain := flag.String("domain", "", "The domain to perform WHOIS lookup on")

	flag.CommandLine.SetOutput(os.Stderr)
	flag.Parse()

	if *domain == "" {
		fmt.Fprintln(os.Stderr, "Error: 'domain' argument is required")
		flag.Usage()
		return
	}

	fmt.Fprintf(os.Stderr, "Performing mock WHOIS lookup for: %s\n", *domain)
}
