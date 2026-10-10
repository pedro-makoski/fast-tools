package converter

import "context"

func convertAudio(src, dst FormatFile) error {
	return runFFmpeg(context.Background(), src.path, dst.path)
}
