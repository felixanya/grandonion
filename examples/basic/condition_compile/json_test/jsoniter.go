//+build jsoniter

package json_test

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func MarshalIdent(v interface{}, prefix, indent string) ([]byte, error) {
	fmt.Println("Use [Jsoniter] package")
	return json.MarshalIndent(v, prefix, indent)
}
