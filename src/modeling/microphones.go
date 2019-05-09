package modeling

type Microphone interface {
	SendData() error
}

type MicrophoneProxy struct{}

func (mp *MicrophoneProxy) SendData() error {
	return nil
}

func GenerateDataForMicrophones() {

}
