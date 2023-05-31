package server

import (
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"regexp"
)

var logRegex = regexp.MustCompile("^(?P<time>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}) \\[(?P<type>[A-Z]+)]\\s+(?P<name>\\S+): (?P<message>[^\\n]*)")

func parseLogLine(logLine string) *proto.Log {
	match := logRegex.FindStringSubmatch(logLine)
	return &proto.Log{
		DateTime: match[1],
		Level:    match[2],
		Name:     match[3],
		Message:  match[4],
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
