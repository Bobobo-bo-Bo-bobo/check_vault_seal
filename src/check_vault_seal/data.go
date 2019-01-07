package main

type VaultSealStatus struct {
	Type                  string  `json:"type"`
	Initialized           bool    `json:"initialized"`
	Sealed                bool    `json:"sealed"`
	RequiredKeysForUnseal int     `json:"t"`
	NumberOfSealKeys      int     `json:"n"`
	Progress              int     `json:"progress"`
	Nonce                 string  `json:"nonce"`
	Version               string  `json:"version"`
	Migration             bool    `json:"migration"`
	RecoverySeal          bool    `json:"recovery_seal"`
	ClusterName           *string `json:"cluster_name"`
	ClusterId             *string `json:"cluster_id"`
	HTTPStatusCode        int
	HTTPStatus            string
}

// Nagios/Icinga exit codes
const (
	NAGIOS_OK int = iota
	NAGIOS_WARNING
	NAGIOS_CRITICAL
	NAGIOS_UNKNOWN
)

type NagiosState struct {
	Critical []string
	Warning  []string
	Ok       []string
	Unknown  []string
	PerfData []string
}
