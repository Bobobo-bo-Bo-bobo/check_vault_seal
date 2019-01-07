package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var vault_addr = flag.String("addr", "https://127.0.0.1:8200", "Vault address")
	var insecure_ssl = flag.Bool("insecure-ssl", false, "Don't validate server SSL certificate")
	var timeout = flag.Int("timeout", 15, "Connection timeout")

	flag.Usage = showUsage

	flag.Parse()

	seal, err := FetchVaultState(vault_addr, *insecure_ssl, *timeout, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(NAGIOS_UNKNOWN)
	}

	seal_state := ParseVaultState(seal)
	rc, msg := ProcessStatus(seal_state)
	fmt.Println(msg)
	os.Exit(rc)
}
