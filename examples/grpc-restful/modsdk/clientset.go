// Code generated by protoc-gen-grpc-gateway-gosdk. DO NOT EDIT.
// versions:
//    protoc-gen-grpc-gateway-gosdk 1.7.1-next-6ba19e
// type: clientset

package autosdk

import (
	"github.com/jaronnie/autosdk/rest"
	"github.com/jaronnie/autosdk/typed"
	userv1 "github.com/jaronnie/autosdk/typed/userv1"
)

type Interface interface {
	Direct() typed.DirectInterface

	Userv1() userv1.Userv1Interface
}

type Clientset struct {
	// direct client to request
	direct *typed.DirectClient

	userv1 *userv1.Userv1Client
}

func (x *Clientset) Direct() typed.DirectInterface {
	return x.direct
}

func (x *Clientset) Userv1() userv1.Userv1Interface {
	return x.userv1
}

func NewClientWithOptions(ops ...rest.Opt) (*Clientset, error) {
	c := &rest.RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	configShallowCopy := *c
	var cs Clientset
	var err error
	cs.direct, err = typed.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.userv1, err = userv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}
