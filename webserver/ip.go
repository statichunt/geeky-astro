package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slog"
)

var url = "http://api-production-netinfo.hyvesdp.com/v2/network?ip_address=197.94.4.80"
var api_header_key = "x-api-key"

type IpService interface {
	GetNetwork(ip string) (*Network, error)
}

type Network struct {
	Type       string            `json:"type"`
	Attributes NetworkAttributes `json:"attributes"`
}

type NetworkAttributes struct {
	Network        string `json:"network"`
	Mcc            string `json:"mcc"`
	Mnc            string `json:"mnc"`
	VerifiedOnnet  bool   `json:"verified_onnet"`
	CountryIsoCode string `json:"country_iso_code"`
	Range          string `json:"range"`
}

type IpServiceResponse struct {
	Data Network `json:"data"`
}

type IpServiceImpl struct {
	logger *slog.Logger
	url    string
	secret string
}

func NewIpService(logger *slog.Logger, url string, secret string) IpService {
	return &IpServiceImpl{logger, url, secret}
}

func (s *IpServiceImpl) GetNetwork(ip string) (*Network, error) {
	// create a new http get request
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?ip_address=%s", s.url, ip), nil)
	if err != nil {
		s.logger.Error("Error creating request", "error", err)
		return nil, err
	}
	req.Header.Add(api_header_key, s.secret)

	// execute the request
	resp, err := httpClient.Do(req)
	if err != nil {
		s.logger.Error("Error executing request", "error", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Error reading response body", "error", err)
		return nil, err
	}

	// unmarshal the response
	var ipServiceResponse IpServiceResponse
	err = json.Unmarshal(body, &ipServiceResponse)
	if err != nil {
		s.logger.Error("Error unmarshalling response", "error", err)
		return nil, err
	}
	s.logger.Debug("Response", "response", ipServiceResponse)
	return &ipServiceResponse.Data, nil
}

func (n *Network) IsOnNet() bool {
	// if mcc or mnc is null return false
	if n.Attributes.Mcc == "" || n.Attributes.Mnc == "" {
		return false
	}
	return true
}

//Offnet example
// {
//     "data": {
//         "type": "network",
//         "attributes": {
//             "network": "OPTINET",
//             "mcc": null,
//             "mnc": null,
//             "verified_onnet": false,
//             "country_iso_code": "ZA",
//             "range": "197.94.0.0/21"
//         }
//     },
//     "meta": {}
// }

//Onnet example
// {
//     "data": {
//         "type": "network",
//         "attributes": {
//             "network": "Telkom Internet",
//             "mcc": "655",
//             "mnc": "02",
//             "verified_onnet": false,
//             "country_iso_code": "ZA",
//             "range": "102.249.4.0/23"
//         }
//     },
//     "meta": {}
// }
