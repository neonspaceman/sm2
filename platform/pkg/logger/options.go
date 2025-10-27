package logger

type Option func(logger *Logger)

func WithDevelopmentMode(developmentMode bool) Option {
	return func(logger *Logger) {
		logger.development = developmentMode
	}
}

func WithLevel(level int8) Option {
	return func(logger *Logger) {
		logger.level = level
	}
}
