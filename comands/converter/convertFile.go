package converter

type FormatFile struct {
	path       string
	typeOfFile int
	foramtFile int
}

type ConvertFile struct {
	src FormatFile
	dst FormatFile
}

func (c *ConvertFile) convert() {

}
