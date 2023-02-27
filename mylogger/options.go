package mylogger

var option = &Option{}
var defaultOptions = []optionFun{
	WhitLevel("info"),
	WhitLoggingDir("logs"),
	WhitMaxSize(10),
	WhitMaxAge(5),
	WhitMaxBackup(5),
	WhitIsCompress(false),
	WhitIsConsole(true),
}

type Option struct {
	Level      string `yaml:"level"`
	LoggingDir string `yaml:"dir"`
	IsCompress bool   `yaml:"is_compress"`
	IsConsole  bool   `yaml:"is_console"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackup  int    `yaml:"max_backup"`
}

type optionFun = func(o *Option)

func WhitLevel(level string) optionFun {
	return func(o *Option) {
		o.Level = level
	}
}

func WhitLoggingDir(dir string) optionFun {
	return func(o *Option) {
		o.LoggingDir = dir
	}
}

func WhitMaxSize(maxSize int) optionFun {
	return func(o *Option) {
		o.MaxSize = maxSize
	}
}

func WhitMaxAge(maxAge int) optionFun {
	return func(o *Option) {
		o.MaxAge = maxAge
	}
}

func WhitMaxBackup(maxBackup int) optionFun {
	return func(o *Option) {
		o.MaxBackup = maxBackup
	}
}

func WhitIsCompress(isCompress bool) optionFun {
	return func(o *Option) {
		o.IsCompress = isCompress
	}
}

func WhitIsConsole(isConsole bool) optionFun {
	return func(o *Option) {
		o.IsConsole = isConsole
	}
}
