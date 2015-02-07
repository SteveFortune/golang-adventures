
package atourofgo


import (
	"fmt"
	"errors"
)


/**
*	Different logging levels
*/
const (
	kLogLevelInfo = iota
	kLogLevelWarning = iota
	kLogLevelError = iota
)


/**
*	Logger interface
*/
type Logger interface {

	Info(log string, params ...string) error
	Warn(log string, params ...string) error
	Err(log string, params ...string) error

}


/**
*	A logger. This is an example which uses `fmt.Printf` to log
*	to the standard output stream.
*
*	I realise that there is already a standard API for dealing with
*	logging in GoLang, but its really just for demonstrational purposes.
*
*	Implements the Logger interface
*/
type StdOutLogger struct {

	/**
	*	A string that the logs are all prefixed with. Defaults to
	* 	'Logger:'
	*
	*	@public
	*/
	LogPrefix string
}


/**
*	Initialise a new StdOutLogger instance
*
*	@public
*/
func NewStdOutLogger() *StdOutLogger {

	logger := &StdOutLogger{
		LogPrefix: "Logger:",
	}

	return logger
}


/**
*	Builds the logging prefix from both the LogPrefix property and
*	the given log level.
*
*	@param 	level	Should be either Info, Warning or Error
*	@return	string	The constructed logging prefix
*	@return error	An error or nil if no error occured
*	@private
*/
func (logger *StdOutLogger) buildPrefix(level int) (string, error){

	logLevelMap := map[int] string{
		kLogLevelInfo: "<I>",
		kLogLevelWarning: "<W>",
		kLogLevelError: "<E>",
	}

	var formattedPrefix string
	var err error

	levelStr, levelValid := logLevelMap[level]

	if !levelValid {
		err = errors.New("Log level is invalid")
	} else {
		formattedPrefix = fmt.Sprintf("%q: %q ...", logger.LogPrefix, levelStr)
	}

	return formattedPrefix, err
}


/**
*	Creates a formatted log string from the given log level.
*/
func (logger *StdOutLogger) log(logLevel int, log string, params []string) error {

	prefix, err := logger.buildPrefix(logLevel)

	if err != nil {
		formattedLog := fmt.Sprintf("%q %q", prefix, log)
		fmt.Printf(formattedLog, params)
	}

	return err
}


/**
*	Logs with the `Info` prefix
*
*	@public
*/
func (logger *StdOutLogger) Info(log string, params ...string) error {

	return logger.log(kLogLevelInfo, log, params);
}


/**
*	Logs with the `Warning` prefix
*
*	@public
*/
func (logger *StdOutLogger) Warn(log string, params ...string) error {

	return logger.log(kLogLevelWarning, log, params);
}


/**
*	Logs with the `Error` prefix
*
*	@public
*/
func (logger *StdOutLogger) Err(log string, params ...string) error {

	return logger.log(kLogLevelError, log, params);
}
