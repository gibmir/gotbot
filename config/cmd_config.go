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

type CmdConfigFactory struct {
	reader *Reader
}

func NewCmdReader() *Reader {
	args := make(map[string]string)
	key := ""
	for _, osArg := range os.Args {
		if strings.Contains(osArg, "-") {
			key = osArg
			args[osArg] = ""
		} else {
			args[key] = osArg
		}
	}
	return &Reader{args}
}

// configuration factory that uses command line arguments
func (factory *CmdConfigFactory) Create() (*Config, error) {
	token := factory.reader.Read(TokenArgName)
	if token == "" {
		return nil, ErrArgEmptyToken
	}
	return &Config{Token: token} /*success*/, nil /*error*/
}
