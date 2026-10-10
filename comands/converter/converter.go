package converter

import (
	"errors"
	"fmt"
	"os"
)

type ConverterConf struct {
	files []ConvertFile
}

func (c *ConverterConf) Do() error {
	comands := os.Args[2:]

	if len(comands) < 2 {
		return errors.New("Insuficient params")
	}

	originMode := comands[0]
	dstMode := comands[1]

	files, err := getToConvertFiles(originMode, dstMode)

	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println("Starting to convert ", file.src.path,
			" to: ", file.dst.path)
		err := file.convert()

		if err != nil {
			return err
		}

		fmt.Println("File ", file.src.path, " converted with sucess to: ", file.dst.path)
	}

	return nil
}
