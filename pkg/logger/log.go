package logger

type Logger interface {
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	InfowFields(msg string, fields map[string]interface{})
	WarnwFields(msg string, fields map[string]interface{})
	ErrorwFields(msg string, fields map[string]interface{})
}

// Regular logger functions
func Infof(format string, args ...interface{}) {
	if log != nil {
		log.Infof(format, args...)
	}
}

func Infoln(args ...interface{}) {
	if log != nil {
		log.Info(args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if log != nil {
		log.Warnf(format, args...)
	}
}

func Warnln(args ...interface{}) {
	if log != nil {
		log.Warn(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if log != nil {
		log.Errorf(format, args...)
	}
}

func Errorln(args ...interface{}) {
	if log != nil {
		log.Error(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if log != nil {
		log.Debugf(format, args...)
	}
}

func Debugln(args ...interface{}) {
	if log != nil {
		log.Debug(args...)
	}
}

// Structured logging functions
func InfowFields(msg string, fields map[string]interface{}) {
	if log != nil {
		log.With(convertMapToFields(fields)...).Info(msg)
	}
}

func WarnwFields(msg string, fields map[string]interface{}) {
	if log != nil {
		log.With(convertMapToFields(fields)...).Warn(msg)
	}
}

func ErrorwFields(msg string, fields map[string]interface{}) {
	if log != nil {
		log.With(convertMapToFields(fields)...).Error(msg)
	}
}

// Helper function to convert map to zap fields
func convertMapToFields(fields map[string]interface{}) []interface{} {
	result := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		result = append(result, k, v)
	}
	return result
}
