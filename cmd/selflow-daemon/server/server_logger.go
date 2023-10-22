package server

import (
	"encoding/json"
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"strings"
)

type logMessage struct {
	Message  string `json:"msg"`
	Level    string `json:"level"`
	DateTime string `json:"time"`
	Name     string `json:"stepId"`
}

func parseLogLine(logLine string) *proto.Log {
	logAsBytes := []byte(logLine)
	lm := logMessage{}
	err := json.Unmarshal(logAsBytes, &lm)
	if err != nil {
		return nil
	}

	return &proto.Log{
		DateTime: lm.DateTime,
		Level:    strings.ToUpper(lm.Level),
		Name:     lm.Name,
		Message:  lm.Message,
		Metadata: logAsBytes,
	}

}

func (s *Server) GetLogStream(request *proto.GetLogStream_Request, stream proto.Daemon_GetLogStreamServer) error {
	ch, err := s.LogFactory.GetRunReader(request.GetRunId())
	if err != nil {
		return err
	}

	for line := range ch {
		err := stream.Send(parseLogLine(line))
		if err != nil {
			return err
		}
	}

	return nil
}
