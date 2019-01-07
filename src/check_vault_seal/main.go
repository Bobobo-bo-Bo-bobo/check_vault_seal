package main

import (
	"flag"
	"fmt"
	"os"
)

const default_addr string = "https://127.0.0.1:8200"

func main() {
	var vault_addr = flag.String("addr", "", "Vault address")
	var insecure_ssl = flag.Bool("insecure-ssl", false, "Don't validate server SSL certificate")

	flag.Usage = showUsage

	flag.Parse()

	if *vault_addr == "" {
		*vault_addr = default_addr
	}

	seal, err := FetchVaultState(vault_addr, *insecure_ssl, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(NAGIOS_UNKNOWN)
	}

	seal_state := ParseVaultState(seal)
	rc, msg := ProcessStatus(seal_state)
	fmt.Println(msg)
	os.Exit(rc)
}
