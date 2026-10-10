package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getToConvertFiles(baseSrc string, baseDst string) ([]ConvertFile, error) {
	srcFiles, err := resolveSourceFiles(baseSrc)
	if err != nil {
		return nil, err
	}

	dstFormat, err := getFormatFromPath(baseDst)
	if err != nil {
		return nil, err
	}

	dstType := getTypeFromFormat(dstFormat)

	dstPath := filepath.FromSlash(baseDst)
	dstDir := filepath.Dir(dstPath)
	dstName := filepath.Base(dstPath)

	hasWildcard := strings.HasPrefix(dstName, "*.")

	res := make([]ConvertFile, 0, len(srcFiles))

	for _, src := range srcFiles {
		var dstPath string

		srcBaseName := filepath.Base(src.path)
		srcName := strings.TrimSuffix(srcBaseName, filepath.Ext(srcBaseName))

		if hasWildcard {
			dstPath = filepath.Join(
				filepath.Dir(src.path),
				srcName+filepath.Ext(baseDst),
			)
		} else if len(srcFiles) == 1 {
			dstPath = filepath.Join(dstDir, dstName)
		} else {
			dstPath = filepath.Join(dstDir, srcName+filepath.Ext(baseDst))
		}

		res = append(res, ConvertFile{
			src: src,
			dst: FormatFile{
				path:       dstPath,
				typeOfFile: dstType,
				foramtFile: dstFormat,
			},
		})
	}

	return res, nil
}

func resolveSourceFiles(value string) ([]FormatFile, error) {
	path := filepath.FromSlash(value)
	format, err := getFormatFromPath(path)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(value, "*.") {
		matches, err := filepath.Glob(path)
		if err != nil {
			return nil, err
		}

		files := make([]FormatFile, 0, len(matches))

		for _, match := range matches {
			info, err := os.Stat(match)
			if err != nil {
				return nil, err
			}

			if info.IsDir() {
				continue
			}

			files = append(files, FormatFile{
				path:       match,
				typeOfFile: getTypeFromFormat(format),
				foramtFile: format,
			})
		}

		return files, nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("arquivo não encontrado: %s", path)
	}

	if info.IsDir() {
		return nil, fmt.Errorf("o caminho é um diretório: %s", path)
	}

	return []FormatFile{
		{
			path:       path,
			typeOfFile: getTypeFromFormat(format),
			foramtFile: format,
		},
	}, nil
}

func getFormatFromPath(path string) (int, error) {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".mp3":
		return MP3, nil
	case ".wav":
		return WAV, nil
	case ".ogg":
		return OGG, nil
	case ".flac":
		return FLAC, nil
	case ".aac":
		return AAC, nil
	case ".m4a":
		return M4A, nil
	case ".mp4":
		return MP4, nil
	case ".mov":
		return MOV, nil
	case ".avi":
		return AVI, nil
	case ".mkv":
		return MKV, nil
	case ".webm":
		return WEBM, nil
	case ".wmv":
		return WMV, nil
	case ".flv":
		return FLV, nil
	case ".png":
		return PNG, nil
	case ".jpg":
		return JPG, nil
	case ".jpeg":
		return JPEG, nil
	case ".webp":
		return WEBP, nil
	case ".gif":
		return GIF, nil
	case ".bmp":
		return BMP, nil
	case ".tiff":
		return TIFF, nil
	case ".ico":
		return ICO, nil
	default:
		return 0, fmt.Errorf("formato não suportado: %s", ext)
	}
}

func getTypeFromFormat(format int) int {
	switch format {
	case MP3, WAV, OGG, FLAC, AAC, M4A:
		return AUDIO
	case MP4, MOV, AVI, MKV, WEBM, WMV, FLV:
		return VIDEO
	case PNG, JPG, JPEG, WEBP, GIF, BMP, TIFF, ICO:
		return IMAGE
	default:
		panic("formato inválido")
	}
}
