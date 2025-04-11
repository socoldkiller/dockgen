package gen_container

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyIO(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		//todo
	}
}

func convertMounts(mounts map[string]string) (map[string]struct{}, []string, error) {
	volumes := make(map[string]struct{})
	binds := make([]string, 0, len(mounts))

	for hostRel, containerPath := range mounts {
		absHostPath, err := filepath.Abs(hostRel)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve absolute path for %s: %w", hostRel, err)
		}

		if _, err := os.Stat(absHostPath); err != nil {
			return nil, nil, fmt.Errorf("host path %s does not exist: %w", absHostPath, err)
		}

		volumes[containerPath] = struct{}{}
		binds = append(binds, fmt.Sprintf("%s:%s", absHostPath, containerPath))
	}

	return volumes, binds, nil
}
