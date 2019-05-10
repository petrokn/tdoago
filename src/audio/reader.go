package audio

type FileReader interface {
	ReadDataFromFile(path *string)
}

type BaseFileReader struct{}

func (br *BaseFileReader) ReadDataFromFile(path *string) {

}
