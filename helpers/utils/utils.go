package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func LogPrint(jsonData interface{}) {
	prettyJSON, _ := json.MarshalIndent(jsonData, "", "")
	fmt.Printf("%s\n", strings.ReplaceAll(string(prettyJSON), "\n", ""))
}

func DebugJson(value interface{}) {
	fmt.Println(reflect.TypeOf(value).String())
	prettyJSON, _ := json.MarshalIndent(value, "", "    ")
	fmt.Printf("%s\n", string(prettyJSON))
}
