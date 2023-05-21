package container

import (
	"io"
	"regexp"
	"sync"
)

var OutputRegex = regexp.MustCompile("^::output::([a-zA-Z][a-zA-Z_]*)::(\\S*)")

type writerWithOutput struct {
	io.Writer
	output     map[string]string
	outputLock sync.Mutex
}

func (w *writerWithOutput) Write(p []byte) (n int, err error) {

	matches := OutputRegex.FindSubmatch(p)
	if len(matches) == 3 {
		outputName := string(matches[1])
		outputValue := string(matches[2])
		w.outputLock.Lock()
		w.output[outputName] = outputValue
		w.outputLock.Unlock()
	}

	return w.Writer.Write(p)
}
