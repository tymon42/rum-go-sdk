package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func newHTTPClient() (*http.Client, error) {
	certPath, err := filepath.Abs("C:\\Users\\why00\\Documents\\project\\go\\src\\2-some-one-cli\\2SOMEone\\certs\\server.crt")

	if err != nil {
		return nil, err
	}

	if certPath != "" {
		caCert, err := ioutil.ReadFile(certPath)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		if err != nil {
			return nil, err
		}
		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}
		// if config.RumConfig.Quorum.ServerSSLInsecure {
			tlsConfig.InsecureSkipVerify = true
		// }

		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig, DisableKeepAlives: true}
		// 5 seconds timeout, all timeout will be ignored, since we refresh all data every half second
		return &http.Client{Transport: transport, Timeout: 5 * time.Second}, nil
	}
	return &http.Client{}, nil
}

func HttpGet(url string) ([]byte, error) {
	client, err := newHTTPClient()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func HttpPost(url string, data []byte) ([]byte, error) {
	client, err := newHTTPClient()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "error") {
		return nil, errors.New(string(body))
	}

	return body, err
}
func HttpDelete(url string, data []byte) ([]byte, error) {
	client, err := newHTTPClient()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
