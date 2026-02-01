package pkg

import (
	"fmt"
	"runtime"
	"strings"
)

const ModuleName = "github.com/emel-study/emel-study-server/"

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", strings.Replace(runtime.FuncForPC(pc).Name(), ModuleName, "", -1))
}
