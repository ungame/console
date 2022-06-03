package console

import (
	"encoding/json"
	"fmt"
)

func Log(data interface{}, indent string) error {
	b, err := json.MarshalIndent(data, "", indent)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
