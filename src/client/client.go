package client

import "tdoago/src/modeling"

type Client interface {
	Send(microphones []*modeling.Microphone) []error
	SendLive(microphones []*modeling.Microphone) []error
}

type BaseClient struct{}

func (bc *BaseClient) Send(microphones []*modeling.Microphone) []error {
	errors := make([]error, len(microphones))

	for _, microphone := range microphones {
		errors = append(errors, (*microphone).(modeling.Microphone).SendData())
	}

	return errors
}

func (bc *BaseClient) SendLive(microphones []*modeling.Microphone) []error {
	return nil
}
