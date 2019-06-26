// +build !jsoniter

package json_test

import (
	"encoding/json"
	"fmt"
)

func MarshalIdent(v interface{}, prefix, indent string) ([]byte, error) {
	fmt.Println("Use [encoding/json] package ")
	return json.MarshalIndent(v, prefix, indent)
}
