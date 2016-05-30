package leftpad

import (
	"bytes"
	"unicode/utf8"
)

func Leftpad(str string, len int, ch string) string {
	strlen := utf8.RuneCountInString(str)

	if len == strlen {
		return str
	}

	if len < strlen {
		return string([]rune(str)[0:len])
	}

	var buffer bytes.Buffer

	diff := len - strlen
	for i := 0; i < diff; i++ {
		buffer.WriteString(ch)
	}

	return buffer.String() + str
}
