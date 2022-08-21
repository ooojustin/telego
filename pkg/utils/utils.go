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

func Contains[T comparable](haystack []T, needle T) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func Remove[T comparable](haystack []T, needle T) []T {
    for idx, item := range haystack {
        if item == needle {
            return append(haystack[:idx], haystack[idx+1:]...)
        }
    }
    return haystack
}
