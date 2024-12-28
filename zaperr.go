package zaperr

import (
	"go.uber.org/zap"
)

type Zaperr struct {
	logger *zap.Logger
}

func NewZaperr(logger *zap.Logger) *Zaperr {
	return &Zaperr{
		logger: logger,
	}
}

func (s *Zaperr) LogError(err error, message string, additionalLogic ...func()) error {
	if err != nil {
		s.logger.Error(message, zap.Error(err))
		if len(additionalLogic) > 0 {
			for _, fn := range additionalLogic {
				fn()
			}
		}
		return err
	}

	return nil
}

func (s *Zaperr) LogPanicError(err error, message string, additionalLogic ...func()) {
	if err != nil {
		s.logger.Panic(message, zap.Error(err))
		if len(additionalLogic) > 0 {
			for _, fn := range additionalLogic {
				fn()
			}
		}
	}
}

func (s *Zaperr) LogWarningError(err error, message string, additionalLogic ...func()) error {
	if err != nil {
		s.logger.Warn(message, zap.Error(err))
		if len(additionalLogic) > 0 {
			for _, fn := range additionalLogic {
				fn()
			}
		}
		return err
	}
	return nil
}

func (s *Zaperr) LogDebugError(err error, message string, additionalLogic ...func()) {
	if err != nil {
		s.logger.Debug(message, zap.Error(err))

		if len(additionalLogic) > 0 {
			for _, fn := range additionalLogic {
				fn()
			}
		}
	}
}

func (s *Zaperr) LogInfoError(err error, message string, additionalLogic ...func()) {
	if err != nil {
		s.logger.Info(message, zap.Error(err))
		if len(additionalLogic) > 0 {
			for _, fn := range additionalLogic {
				fn()
			}
		}
	}
}
