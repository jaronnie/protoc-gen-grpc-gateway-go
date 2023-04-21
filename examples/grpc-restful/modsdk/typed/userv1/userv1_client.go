// Code generated by protoc-gen-grpc-gateway-gosdk. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-gosdk 1.7.1-next-6ba19e
// type: userv1_client

package userv1

import (
	"github.com/jaronnie/autosdk/rest"
)

type Userv1Interface interface {
	RESTClient() rest.Interface

	UserGetter
}

type Userv1Client struct {
	restClient rest.Interface
}

func (x *Userv1Client) RESTClient() rest.Interface {
	if x == nil {
		return nil
	}
	return x.restClient
}

func (x *Userv1Client) User() UserInterface {
	return newUserClient(x)
}

// NewForConfig creates a new Userv1Client for the given config.
func NewForConfig(x *rest.RESTClient) (*Userv1Client, error) {
	config := *x
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Userv1Client{client}, nil
}
