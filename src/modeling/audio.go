package modeling

type AudioFileReader interface {
	ReadDataFromFile(path *string)
}

type BaseAudioFileReader struct {}

func (br *BaseAudioFileReader) ReadDataFromFile(path *string) {

}
