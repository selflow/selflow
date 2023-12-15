package workflow

type Status struct {
	code        uint
	name        string
	finished    bool
	cancellable bool
	executable  bool
}

func BuildStatus(code uint, name string, finished bool, cancellable bool, executable bool) Status {
	return Status{
		code:        code,
		name:        name,
		finished:    finished,
		cancellable: cancellable,
		executable:  executable,
	}
}

func (s Status) GetCode() uint {
	return s.code
}

func (s Status) GetName() string {
	return s.name
}

func (s Status) IsFinished() bool {
	return s.finished
}

func (s Status) IsCancellable() bool {
	return s.cancellable
}
func (s Status) IsExecutable() bool {
	return s.executable
}

const (
	successCode uint = iota
	errorCode
	cancelledCode
	runningCode
	pendingCode
	readyCode
)

func StatusEquals(a, b Status) bool {
	return a.GetCode() == b.GetCode()
}

var (
	//PENDING Step is waiting for its launch
	PENDING = Status{pendingCode, "PENDING", false, true, true}

	//READY Step is going to be launched
	READY = Status{readyCode, "PENDING", false, true, false}

	// RUNNING Step is executing
	RUNNING = Status{runningCode, "RUNNING", false, true, false}

	// SUCCESS Step terminated with no errors
	SUCCESS = Status{successCode, "SUCCESS", true, false, false}

	// ERROR Step terminated with a blocking error
	ERROR = Status{errorCode, "ERROR", true, false, false}

	// CANCELLED Step execution has been interrupted or a requirement exited with a blocking error
	CANCELLED = Status{cancelledCode, "CANCELLED", true, false, false}
)
