
package atourofgo


import (
	"fmt"
	"errors"
)


/**
*	Different logging levels
*/
const (
	Info = iota
	Warning = iota
	Error = iota
)


/**
*	A logger. This is an example which uses `fmt.Printf` to log
*	to the standard output stream.
*
*	I realise that there is already a standard API for dealing with
*	logging in GoLang, but its really just for demonstrational purposes.
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
		logPrefix: "Logger:"
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
func (logger *StdOutLogger) buildPrefix(level int) (prefix string, err error){

	logLevelMap = map[int] string{
		Info: "<I>",
		Warning: "<W>",
		Error: "<E>",
	}

	formattedPrefix string
	err error

	levelStr, levelValid := logLevelMap

	if !levelValid{
		err = errors.New("Log level is invalid")
	}
	else{
		formattedPrefix = fmt.Fprintf("%q: %q ...", logger.LogPrefix, levelStr)
	}

	return formattedPrefix, err
}


/**
*	Logs with the 'Info' prefix
*
*	@public
*/
func (logger *StdOutLogger) LogInfo() {

}


/**
*	Logs with the 'Warning' prefix
*
*	@public
*/
func (logger *StdOutLogger) LogWarn() {

}


/**
*	Logs with the 'Error' prefix
*
*	@public
*/
func (logger *StdOutLogger) LogErr() {

}
