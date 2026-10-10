package converter

type FormatFile struct {
	path       string
	typeOfFile int
	formatfile int
}

type ConvertFile struct {
	src FormatFile
	dst FormatFile
}

func (f FormatFile) Type() (int, bool) {
	switch f.formatfile {
	case MP3, WAV, OGG, FLAC, AAC, M4A:
		return AUDIO, true
	case MP4, MOV, AVI, MKV, WEBM, WMV, FLV:
		return VIDEO, true
	case PNG, JPG, JPEG, WEBP, GIF, BMP, TIFF, ICO:
		return IMAGE, true
	default:
		return 0, false
	}
}
