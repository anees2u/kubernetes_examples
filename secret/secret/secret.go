package secret

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	AwsRegionalDns = "REGIONAL_DNS"
	AwsDnsRegionA  = "DNS_REGION_A"
	AwsDnsRegionB  = "DNS_REGION_B"
	AwsDnsRegionC  = "DNS_REGION_C"
)

type k8sEndpointSecret struct {
	secretName string
	namespace  string
	client     client.Client
}

func New(name string, namespace string, client client.Client) IEndpointSecret {
	return &k8sEndpointSecret{name, namespace, client}
}

func (k *k8sEndpointSecret) CreateEndpointSecret(dnsEntries []types.DnsEntry, region string) (*v1.Secret, error) {
	data := make(map[string][]byte)
	for _, dnsentry := range dnsEntries {
		if strings.Contains(*dnsentry.DnsName, region+"a") {
			b64.StdEncoding.Encode(data[AwsDnsRegionA], []byte(*dnsentry.DnsName))
		} else if strings.Contains(*dnsentry.DnsName, region+"b") {
			b64.StdEncoding.Encode(data[AwsDnsRegionB], []byte(*dnsentry.DnsName))
		} else if strings.Contains(*dnsentry.DnsName, region+"c") {
			b64.StdEncoding.Encode(data[AwsDnsRegionC], []byte(*dnsentry.DnsName))
		} else if strings.Contains(*dnsentry.DnsName, region) {
			b64.StdEncoding.Encode(data[AwsRegionalDns], []byte(*dnsentry.DnsName))
		}

	}
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: k.namespace,
			Name:      k.secretName,
		},
		Data: data,
	}

	err := k.client.Create(context.TODO(), secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (k *k8sEndpointSecret) SetEndpointSecret(dnsEntries []types.DnsEntry, region string) (*v1.Secret, error) {
	data := make(map[string][]byte)
	for _, dnsentry := range dnsEntries {
		if strings.Contains(*dnsentry.DnsName, region+"a") {
			b64.StdEncoding.Encode(data[AwsDnsRegionA], []byte(*dnsentry.DnsName))
		} else if strings.Contains(*dnsentry.DnsName, region+"b") {
			b64.StdEncoding.Encode(data[AwsDnsRegionB], []byte(*dnsentry.DnsName))
		} else if strings.Contains(*dnsentry.DnsName, region+"c") {
			b64.StdEncoding.Encode(data[AwsDnsRegionC], []byte(*dnsentry.DnsName))
		} else if strings.Contains(*dnsentry.DnsName, region) {
			b64.StdEncoding.Encode(data[AwsRegionalDns], []byte(*dnsentry.DnsName))
		}

	}
	fmt.Println(data)

	return nil, nil
}
