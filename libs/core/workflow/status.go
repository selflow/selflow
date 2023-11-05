package workflow

type Status interface {
	GetCode() uint
	GetName() string
	IsFinished() bool
	IsCancellable() bool
}

type SimpleStatus struct {
	Code        uint
	Name        string
	Finished    bool
	Cancellable bool
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

const (
	SuccessCode uint = iota
	ErrorCode
	CancelledCode
	RunningCode
	PendingCode
	InitializingCode
	CreatedCode
)

var (
	SUCCESS      = SimpleStatus{SuccessCode, "SUCCESS", true, false}
	ERROR        = SimpleStatus{ErrorCode, "ERROR", true, false}
	CANCELLED    = SimpleStatus{CancelledCode, "CANCELLED", true, false}
	RUNNING      = SimpleStatus{RunningCode, "RUNNING", false, true}
	PENDING      = SimpleStatus{PendingCode, "PENDING", false, true}
	INITIALIZING = SimpleStatus{InitializingCode, "INITIALIZING", false, true}
	CREATED      = SimpleStatus{CreatedCode, "CREATED", false, true}
)
