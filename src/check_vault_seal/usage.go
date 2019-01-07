package main

import "fmt"

const help string = `check_vault_seal version 1.0.0
Copyright (C) by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

redfish-tool is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Usage: check_vault_seal --addr <vault_address>

	--addr <vault_address>		Address for accessing the Vault API.
								Default: https://127.0.0.1:8200

`

func showUsage() {
	fmt.Println(help)
}
