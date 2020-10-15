// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/tnozicka/openshift-acme/pkg/client/acme/clientset/versioned/typed/acme/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAcmeV1 struct {
	*testing.Fake
}

func (c *FakeAcmeV1) ACMEControllers() v1.ACMEControllerInterface {
	return &FakeACMEControllers{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAcmeV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}