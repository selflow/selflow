package server

import (
	"encoding/json"
	proto2 "github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"strings"
)

type logMessage struct {
	Message  string `json:"msg"`
	Level    string `json:"level"`
	DateTime string `json:"time"`
	Name     string `json:"stepId"`
}

func parseLogLine(logLine string) *proto2.Log {
	logAsBytes := []byte(logLine)
	lm := logMessage{}
	err := json.Unmarshal(logAsBytes, &lm)
	if err != nil {
		return nil
	}

	return &proto2.Log{
		DateTime: lm.DateTime,
		Level:    strings.ToUpper(lm.Level),
		Name:     lm.Name,
		Message:  lm.Message,
		Metadata: logAsBytes,
	}

}

func (s *Server) GetLogStream(request *proto2.GetLogStream_Request, stream proto2.Daemon_GetLogStreamServer) error {
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
