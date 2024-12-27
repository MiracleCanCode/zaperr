package zaperr

import "go.uber.org/zap"

func LogError(logger *zap.Logger, err error, message string) error {
	if err != nil {
		if message == "" {
			message = "Unknown error"
		}
		logger.Error(message, zap.Error(err))
		return err
	}
	return nil
}

func LogPanicError(logger *zap.Logger, err error, message string) {
	if err != nil {
		if message == "" {
			message = "Unknown panic"
		}
		logger.Panic(message, zap.Error(err))
	}
}

func LogWarningError(logger *zap.Logger, err error, message string) error {
	if err != nil {
		if message == "" {
			message = "Unknown warning"
		}
		logger.Warn(message, zap.Error(err))
		return err
	}
	return nil
}
