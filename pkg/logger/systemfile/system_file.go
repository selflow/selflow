package systemfile

import (
	"fmt"
	"io"
	"os"
	"path"
)

type LogFactory struct {
	BaseDirectory string
}

func (l *LogFactory) GetRunLogger(runId string) (io.Reader, io.WriteCloser, error) {
	file, err := os.Create(path.Join(l.BaseDirectory, fmt.Sprintf("run-%s.log", runId)))

	return file, file, err
}
