package config

import (
	"fmt"
	"os"
	"strings"
)

/*command line args*/
const TokenArgName = "-t"

var ErrArgEmptyToken = fmt.Errorf("[%v] argument wasn't specified",
	TokenArgName)

type CmdReader struct {
	args map[string]string
}

type CmdConfigFactory struct {
	reader CmdReader
}

func CreateReader() *CmdReader {
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
	return &CmdReader{args}
}

func (reader CmdReader) Read(key string) string {
	return reader.args[key]
}

func (reader CmdReader) DefaultRead(key, defaultVal string) string {
	val := reader.args[key]
	if val == "" {
		return defaultVal
	}
	return val
}

// configuration factory that uses command line arguments
func (factory CmdConfigFactory) Create() (*Config, error) {
	token := factory.reader.Read(TokenArgName)
	if token == "" {
		return nil, ErrArgEmptyToken
	}
	return &Config{Token: token} /*success*/, nil /*error*/
}
