package bug

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/tal-tech/go-zero/tools/goctl/internal/version"
)

type env map[string]string

func (e env) string() string {
	if e == nil {
		return ""
	}

	w := bytes.NewBuffer(nil)
	for k, v := range e {
		w.WriteString(fmt.Sprintf("%s = %q\n", k, v))
	}

	return w.String()
}

func getEnv() env {
	e := make(env)
	e[os] = runtime.GOOS
	e[arch] = runtime.GOARCH
	e[goctlVersion] = version.BuildVersion
	return e
}
