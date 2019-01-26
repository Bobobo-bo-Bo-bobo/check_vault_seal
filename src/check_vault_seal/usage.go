package main

import "fmt"

const help string = `check_vault_seal version 1.0.2
Copyright (C) by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

check_vault_seal is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Usage: check_vault_seal --addr <vault_address>

    --addr <vault_address>  Address for accessing the Vault API.
                            Default: https://127.0.0.1:8200

    --ca-cert <file>        Use <file> as CA certificate for servers certificate

    --insecure-ssl          Don't validate server certificate.

    --timeout <timeout>     Connection timeout in seconds.
                            Default: 15

`

func showUsage() {
	fmt.Println(help)
}
