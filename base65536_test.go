package base65536

import (
	"fmt"
	"testing"
)

var testValuesString = []string{
	"hello world",
	"what up",
	"fuck php with a cactus",
	"おれわおちんちんだいすきなんだよ",
	"膿",
	"持能够宠",
	"😀🙎🏼👰🏼🙀",
	"케이팝 놀라운 빌어 먹을된다",
}

func TestMarshal(t *testing.T) {
	for _, val := range testValuesString {
		fmt.Printf(`==> "%s"%s`, val, "\n")
		res := Marshal([]byte(val))
		fmt.Printf(`... "%s"%s`, res, "\n")
	}
}

func TestUnmarshal(t *testing.T) {
	converted := []string{}
	for _, val := range testValuesString {
		converted = append(converted, Marshal([]byte(val)))
	}
	for index, val := range converted {
		out := []byte{}
		err := Unmarshal([]byte(val), &out)
		if err != nil {
			t.Errorf("Error while unmarshaling: %s\n", err)
			continue
		}
		fmt.Println(testValuesString[index], "==>", val, "==>", string(out))
		if testValuesString[index] != string(out) {
			t.Error("(check above log) original value doesn't match end value")
		}
	}
}
