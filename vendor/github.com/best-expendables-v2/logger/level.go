package logger

type Level uint32

const (
	EmergencyLevel Level = iota
	AlertLevel
	CriticalLevel
	ErrorLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
)

func (lvl Level) String() string {
	switch lvl {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case NoticeLevel:
		return "notice"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "err"
	case AlertLevel:
		return "alert"
	case CriticalLevel:
		return "crit"
	case EmergencyLevel:
		return "emerg"
	}

	return "unknown"
}

var levels = []Level{
	EmergencyLevel,
	AlertLevel,
	CriticalLevel,
	ErrorLevel,
	WarningLevel,
	NoticeLevel,
	InfoLevel,
	DebugLevel,
}
