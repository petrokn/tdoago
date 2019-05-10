package modeling

type DataGenerator interface {
	GenerateSourcePosition()
	GenerateMicrophoneData()
	GenerateMicrophonePositions()
}

type BaseGenerator struct{}

func (bs *BaseGenerator) GenerateMicrophoneData() {

}

func (bs *BaseGenerator) GenerateMicrophonePositions() {

}

func (bs *BaseGenerator) GenerateSourcePosition() {

}
