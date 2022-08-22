package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
    "runtime"
	"strings"
)

func GetPrettyJSON(o interface{}) (string, error) {
	obytes, err := json.MarshalIndent(o, "", "	")
	if err != nil {
		return "", err
	}
	return string(obytes), nil
}

func PrettyPrint(o interface{}) {
	str, err := GetPrettyJSON(o)
	if err == nil {
		fmt.Println(str)
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

func GetFunctionName(i interface{}) string {
	nameFull := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(nameFull, ".")
	return parts[len(parts) - 1]
}
