package workflow

type Status interface {
	GetCode() uint
	GetName() string
	IsFinished() bool
	IsCancellable() bool
	IsExecutable() bool
}

type SimpleStatus struct {
	Code        uint
	Name        string
	Finished    bool
	Cancellable bool
	Executable  bool
}

func (s SimpleStatus) GetCode() uint {
	return s.Code
}

func (s SimpleStatus) GetName() string {
	return s.Name
}

func (s SimpleStatus) IsFinished() bool {
	return s.Finished
}

func (s SimpleStatus) IsCancellable() bool {
	return s.Cancellable
}
func (s SimpleStatus) IsExecutable() bool {
	return s.Executable
}

const (
	SuccessCode uint = iota
	ErrorCode
	CancelledCode
	RunningCode
	PendingCode
	InitializingCode
	CreatedCode
	ReadyCode
)

func StatusEquals(a, b Status) bool {
	return a.GetCode() == b.GetCode()
}

var (
	// SUCCESS Step terminated with no errors
	SUCCESS = SimpleStatus{SuccessCode, "SUCCESS", true, false, false}

	// ERROR Step terminated with a blocking error
	ERROR = SimpleStatus{ErrorCode, "ERROR", true, false, false}

	// CANCELLED Step execution has been interrupted or a requirement exited with a blocking error
	CANCELLED = SimpleStatus{CancelledCode, "CANCELLED", true, false, false}

	// RUNNING Step execution is in progress
	RUNNING = SimpleStatus{RunningCode, "RUNNING", false, true, false}

	//PENDING Step is ready to be launch
	PENDING = SimpleStatus{PendingCode, "PENDING", false, true, true}

	//INITIALIZING Step is warming up
	INITIALIZING = SimpleStatus{InitializingCode, "INITIALIZING", false, true, false}

	//CREATED Step is created and can be executed
	CREATED = SimpleStatus{CreatedCode, "CREATED", false, true, true}

	//READY Step is going to be launched
	READY = SimpleStatus{ReadyCode, "CREATED", false, true, false}
)
