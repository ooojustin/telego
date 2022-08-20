package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrettyPrint(o interface{}) {
	obytes, err := json.MarshalIndent(o, "", "	")
	if err == nil {
		ostr := string(obytes)
		fmt.Println(ostr)
	} else {
		fmt.Printf("Failed to prettyPrint: %s\n", err)
	}
}

func Exitf(code int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(code)
}
