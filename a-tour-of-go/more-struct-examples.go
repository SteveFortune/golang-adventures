
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
		Info: "<I>",
		Warning: "<W>",
		Error: "<E>",
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
func (logger *StdOutLogger) formattedLogStr(log string, logLevel int) (formattedLog string, err error) {
	formattedLog, err = logger.buildPrefix(logLevel)
	return
}


/**
*	Logs with the 'Info' prefix
*
*	@public
*/
func (logger *StdOutLogger) LogInfo(log string, params ...string) error {

	formattedLog, err := logger.formattedLogStr(log, Info)

	if err != nil {
		fmt.Printf(formattedLog, params)
	}

	return err
}


/**
*	Logs with the 'Warning' prefix
*
*	@public
*/
func (logger *StdOutLogger) LogWarn(log string, params ...string) error {

	formattedLog, err := logger.formattedLogStr(log, Warning)

	if err != nil {
		fmt.Printf(formattedLog, params)
	}

	return err
}


/**
*	Logs with the 'Error' prefix
*
*	@public
*/
func (logger *StdOutLogger) LogErr(log string, params ...string) error {

	formattedLog, err := logger.formattedLogStr(log, Error)

	if err != nil {
		fmt.Printf(formattedLog, params)
	}

	return err
}
