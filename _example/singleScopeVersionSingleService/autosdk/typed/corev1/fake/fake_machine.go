// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"github.com/jaronnie/autosdk/pb/corev1"
	"github.com/jaronnie/autosdk/rest"
)

var (
	FakeReturnInitMachine     = &rest.Request{}
	FakeReturnDownloadMachine = &rest.Request{}
)

type MachineGetter interface {
	Machine() MachineInterface
}

type MachineInterface interface {
	InitMachine(param *corev1.Machine) (*rest.Request, error)
	DownloadMachine(param *corev1.Machine) (*rest.Request, error)
}

type FakeMachine struct {
	Fake *FakeCorev1
}

func (f *FakeMachine) InitMachine(param *corev1.Machine) (*rest.Request, error) {
	return FakeReturnInitMachine, nil
}

func (f *FakeMachine) DownloadMachine(param *corev1.Machine) (*rest.Request, error) {
	return FakeReturnDownloadMachine, nil
}
