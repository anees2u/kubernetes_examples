package secret

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type EndpointSecret struct {
	Region             string `json:"region,omitempty"`
	AwsAccessKeyID     string `json:"awsAccessKeyID,omitempty"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey,omitempty"`
}

type IEndpointSecret interface {
	SetEndpointSecret(dnsEntries []types.DnsEntry, region string) (*v1.Secret, error)
}
