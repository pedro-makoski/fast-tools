package converter

import (
	"context"
	"os/exec"
)

func runFFmpeg(ctx context.Context, src, dst string) error {
	cmd := exec.CommandContext(ctx, "ffmpeg",
		"-nostdin",
		"-v", "error",
		"-y",
		"-i", src,
		dst,
	)

	return cmd.Run()
}
