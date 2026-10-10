package converter

import "context"

func convertVideo(src, dst FormatFile) error {
	return runFFmpeg(context.Background(), src.path, dst.path)
}
