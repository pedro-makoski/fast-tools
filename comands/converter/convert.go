package converter

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrInvalidFormat         = errors.New("invalid format")
	ErrIncompatibleFormats   = errors.New("incompatible formats")
	ErrUnsupportedConversion = errors.New("unsupported conversion")
)

func (c *ConvertFile) convert() error {
	srcType, srcOK := c.src.Type()
	dstType, dstOK := c.dst.Type()

	if len(filepath.Dir(c.dst.path)) > 0 {
		os.MkdirAll(filepath.Dir(c.dst.path), 0755)
	}

	if !srcOK || !dstOK {
		return ErrInvalidFormat
	}

	if srcType != dstType {
		return ErrIncompatibleFormats
	}

	switch srcType {
	case IMAGE:
		return convertImage(c.src, c.dst)
	case AUDIO:
		return convertAudio(c.src, c.dst)
	case VIDEO:
		return convertVideo(c.src, c.dst)
	default:
		return ErrUnsupportedConversion
	}
}
