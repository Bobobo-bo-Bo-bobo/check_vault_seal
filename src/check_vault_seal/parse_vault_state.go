package main

import (
	"fmt"
)

func ParseVaultState(v VaultSealStatus) NagiosState {
	var perf_str string
	var max_str string
	var min_str = "0"
	var state = NagiosState{
		Critical: make([]string, 0),
		Warning:  make([]string, 0),
		Ok:       make([]string, 0),
		Unknown:  make([]string, 0),
		PerfData: make([]string, 0),
	}

	if !v.Initialized {
		state.Unknown = append(state.Unknown, "Vault is not initialised")
		max_str = "0"
		perf_str, _ = MakePerfDataString("unseal_keys", "0", nil, nil, nil, &min_str, &max_str)
		state.PerfData = append(state.PerfData, perf_str)
		return state
	}

	if v.Sealed {
		if v.Progress == 0 {
			state.Critical = append(state.Critical, "Vault is sealed")
		} else {
			state.Warning = append(state.Warning, fmt.Sprintf("Vault is sealed, unsealing is in progress with %d required keys missing (unseal at %d of %d keys)", v.RequiredKeysForUnseal-v.Progress, v.RequiredKeysForUnseal, v.NumberOfSealKeys))
		}
	} else {
		state.Ok = append(state.Ok, "Vault is unsealed")
	}

	perf_str, _ = MakePerfDataString("unseal_keys", string(v.Progress), nil, nil, nil, &min_str, &max_str)
	state.PerfData = append(state.PerfData, perf_str)
	return state
}
