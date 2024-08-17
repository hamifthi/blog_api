package logger

import "go.uber.org/zap"

type ZapLogger struct {
	log *zap.SugaredLogger
}

func NewLogger(config zap.Config) (*ZapLogger, error) {
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	return &ZapLogger{sugar}, nil
}

func (l *ZapLogger) Info(message string, values ...any) {
	l.log.Infow(message, values...)
}

func (l *ZapLogger) Warn(message string, values ...any) {
	l.log.Warnw(message, values...)
}

func (l *ZapLogger) Error(message string, values ...any) {
	l.log.Errorw(message, values...)
}

func (l *ZapLogger) Debug(message string, values ...any) {
	l.log.Debugw(message, values...)
}
