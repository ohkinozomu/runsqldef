package main

import "C"
import (
	"fmt"
	"os"

	"guthub.com/ohkinozomu/runsqldef/sqldefwrapper"
)

//export Run
func Run(dbType, dbName, user, password, host, portStr, schemaFile *C.char, enableDropTable C.int) C.int {
	err := sqldefwrapper.RunCore(
		C.GoString(dbType),
		C.GoString(dbName),
		C.GoString(user),
		C.GoString(password),
		C.GoString(host),
		C.GoString(portStr),
		C.GoString(schemaFile),
		enableDropTable != 0,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}
	return 0
}

func main() {}
