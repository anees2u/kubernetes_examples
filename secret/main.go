/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

//"k8s.io/client-go/rest"

//
// Uncomment to load all auth plugins
// _ "k8s.io/client-go/plugin/pkg/client/auth"
//
// Or uncomment to load specific auth plugins
// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

import (
	b64 "encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	v1 "k8s.io/api/core/v1"
)

type DnsEntry struct {
	DnsName      string
	HostedZoneId string
}

const (
	AwsRegionalDns = "REGIONAL_DNS"
	AwsDnsRegionA  = "DNS_REGION_A"
	AwsDnsRegionB  = "DNS_REGION_B"
	AwsDnsRegionC  = "DNS_REGION_C"
)

var (
	DnsName      = "vpce-039c55c5c768a12f2-40jpjpcs.vpce-svc-04411852a623456a3.eu-central-1.vpce.amazonaws.com"
	HostedZoneId = "Z273ZU8SZ5RJPC"
	DnsNameA     = "vpce-039c55c5c768a12f2-40jpjpcs-eu-central-1a.vpce-svc-04411852a623456a3.eu-central-1.vpce.amazonaws.com"
	DnsNameB     = "vpce-039c55c5c768a12f2-40jpjpcs-eu-central-1b.vpce-svc-04411852a623456a3.eu-central-1.vpce.amazonaws.com"
)

func main() {
	dnsEntries := []types.DnsEntry{

		{
			DnsName:      &DnsName,
			HostedZoneId: &HostedZoneId,
		},
		{
			DnsName:      &DnsNameA,
			HostedZoneId: &HostedZoneId,
		},
		{
			DnsName:      &DnsNameB,
			HostedZoneId: &HostedZoneId,
		},
	}
	SetEndpointSecret(dnsEntries, "eu-central-1")

}

func SetEndpointSecret(dnsEntries []types.DnsEntry, region string) (*v1.Secret, error) {
	data := make(map[string][]byte)
	for _, dnsentry := range dnsEntries {
		if strings.Contains(*dnsentry.DnsName, region+"a") {
			data[AwsDnsRegionA] = []byte(b64.StdEncoding.EncodeToString([]byte(*dnsentry.DnsName)))
		} else if strings.Contains(*dnsentry.DnsName, region+"b") {
			data[AwsDnsRegionB] = []byte(b64.StdEncoding.EncodeToString([]byte(*dnsentry.DnsName)))
		} else if strings.Contains(*dnsentry.DnsName, region+"c") {
			data[AwsDnsRegionC] = []byte(b64.StdEncoding.EncodeToString([]byte(*dnsentry.DnsName)))
		} else if strings.Contains(*dnsentry.DnsName, region) {
			data[AwsRegionalDns] = []byte(b64.StdEncoding.EncodeToString([]byte(*dnsentry.DnsName)))
		}

	}
	fmt.Println(data)

	return nil, nil
}
