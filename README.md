**_Note:_** Because I'm running my own servers for serveral years, main development is done at at https://git.ypbind.de/cgit/check_vault_seal/

----

# Nagios check to validate the sealing state of a Vault
## Preface
[Hashicorp Vault](https://www.vaultproject.io) provides a secure mechanism for storing private data. Upon a server start
the data vault is sealed, preventing access to the data stored within. Unsealing requires a defined number of keys.

This Nagios check reports the sealing status of a vault.

## Build requirements
### Go!

Obviously because the exporter has been written in [Go!](https://golang.org).

## Command line parameters
Accepted command line parameters of `check_vault_seal`:

* `--addr <vault_address>` - Address for accessing the Vault API (Default: https://127.0.0.1:8200)
* `--ca-cert <file>` - File containing the CA certificate (in PEM format) of the CA signing the servers certificate. Not required if CA is already present in the systems store of trusted CAs.
* `--insecure-ssl` - Don't verify server certificate.
* `--timeout <timeout>` - Connection timeout in seconds. (Default: 15)

## Reported states

* OK - Vault is unsealed
* WARNING - Vault is sealed but unsealing is in progress
* CRTICAL - Vault is sealed, unsealing has not started
* UNKNOWN - Vault is not initialized or any other error has occured.

## License
This program is licenses under [GLPv3](http://www.gnu.org/copyleft/gpl.html).

