package converter

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	_ "image/gif"
)

func convertImage(src, dst FormatFile) error {
	input, err := os.Open(src.path)
	if err != nil {
		return err
	}
	defer input.Close()

	img, _, err := image.Decode(input)
	if err != nil {
		return err
	}

	output, err := os.Create(dst.path)
	if err != nil {
		return err
	}
	defer output.Close()

	switch dst.formatfile {
	case PNG:
		return png.Encode(output, img)
	case JPG, JPEG:
		return jpeg.Encode(output, img, &jpeg.Options{Quality: 85})
	default:
		return errors.New("unsupported image conversion")
	}
}
