package systemfile

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/selflow/selflow/pkg/selflow"
	"io"
	"os"
	"path"
)

type LogFactory struct {
	BaseDirectory string
}

func NewLogFactory(BaseDirectory string) *LogFactory {
	return &LogFactory{
		BaseDirectory: BaseDirectory,
	}
}

func (l *LogFactory) getLogFilename(runId string) string {
	return path.Join(l.BaseDirectory, fmt.Sprintf("run-%s.log", runId))
}

func (l *LogFactory) GetRunLogger(runId string) (io.Reader, io.WriteCloser, error) {
	err := os.MkdirAll(l.BaseDirectory, 0777)
	if err != nil {
		return nil, nil, err
	}
	file, err := os.Create(l.getLogFilename(runId))

	return file, file, err
}

func (l *LogFactory) GetRunReader(runId string) (chan string, error) {
	err := os.MkdirAll(l.BaseDirectory, 0777)
	if err != nil {
		return nil, err
	}

	ch := make(chan string)

	tf, err := tail.TailFile(l.getLogFilename(runId), tail.Config{Follow: true})
	if err != nil {
		return nil, err
	}

	go func() {
		for line := range tf.Lines {
			if line.Text == selflow.TerminationLogText {
				close(ch)
				break
			}
			ch <- line.Text
		}
	}()

	return ch, nil
}
