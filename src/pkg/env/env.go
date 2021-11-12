package env

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func EnvBuildpack(prefix, output string) error {
	var err error
	var file *os.File

	if _, err = os.Stat(output); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(output, os.ModePerm); err != nil {
			return err
		}
	}

	for _, element := range os.Environ() {
		key := strings.Split(element, "=")[0]
		if strings.HasPrefix(key, prefix) {
			fileName := strings.TrimPrefix(key, prefix)
			if file, err = os.OpenFile(fmt.Sprintf("%s/%s", output, fileName), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
				return err
			}
			defer file.Close()
			file.WriteString(os.Getenv(key))
		}
	}
	return err
}
