package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const UserAgent string = "check_vault_seal/1.0.1"

func FetchVaultState(vault_url *string, insecure bool, timeout int, ca *string) (VaultSealStatus, error) {
	var seal VaultSealStatus
	var t *http.Transport
	var parsed *url.URL
	var reader io.Reader
	var raw []byte
	var state_url string
	var certpool *x509.CertPool

	parsed, err := url.Parse(*vault_url)
	if err != nil {
		return seal, err
	}

	if parsed.Scheme == "https" {
		if ca != nil && *ca != "" {
			cacert, err := ioutil.ReadFile(*ca)
			if err != nil {
				return seal, err
			}
			certpool = x509.NewCertPool()
			certpool.AppendCertsFromPEM(cacert)
		}

		t = &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            certpool,
				InsecureSkipVerify: insecure,
			},
		}
	} else if parsed.Scheme == "http" {
		t = &http.Transport{}
	} else {
		return seal, errors.New(fmt.Sprintf("Unsupported URL scheme %s", parsed.Scheme))
	}

	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: t,
	}

	state_url = *vault_url + "/v1/sys/seal-status"
	request, err := http.NewRequest("GET", state_url, reader)
	if err != nil {
		return seal, err
	}

	request.Header.Set("User-Agent", UserAgent)
	response, err := client.Do(request)

	if err != nil {
		return seal, err
	}

	defer response.Body.Close()

	raw, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return seal, err
	}

	seal.HTTPStatusCode = response.StatusCode
	seal.HTTPStatus = response.Status

	if response.StatusCode != 200 {
		return seal, errors.New(fmt.Sprintf("HTTP request to %s returned %s\n", state_url, response.Status))
	}

	err = json.Unmarshal(raw, &seal)
	if err != nil {
		fmt.Printf("Raw: %s\n", string(raw))
		return seal, err
	}

	return seal, nil
}
