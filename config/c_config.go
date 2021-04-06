package config

import (
	"fmt"
	"os"
	"strings"
)

/*command line args*/
const tokenArgName = "-t"

var ErrArgEmptyToken = fmt.Errorf("[%v] argument is null", tokenArgName)

type CommandLineReader struct {
	args map[string]string
}

type CommandLineConfigFactory struct {
	reader CommandLineReader
}

func Create() *CommandLineReader {
	args := make(map[string]string)
	osArgs := os.Args
	key := ""
	for _, osArg := range osArgs {
		if strings.Contains(osArg, "-") {
			key = osArg
			args[osArg] = ""
		} else {
			args[key] = osArg
		}
	}
	return &CommandLineReader{args}
}

func (reader CommandLineReader) Read(key string) string {
	return reader.args[key]
}

func (reader CommandLineReader) DefaultRead(key, defaultVal string) string {
	val := reader.args[key]
	if val == "" {
		return defaultVal
	}
	return val
}

// configuration factory that uses command line arguments
func (factory CommandLineConfigFactory) Create() (*Config, error) {
	token := factory.reader.Read(tokenArgName)
	if token == "" {
		return nil, ErrArgEmptyToken
	}
	return &Config{Token: token} /*success*/, nil /*error*/
}
