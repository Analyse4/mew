package stringbuilder

import "fmt"

const maxlength = 1000

type stringBuilder struct {
	buffer [maxlength]rune
	offset int
}

// New initialize and return stringbuilder object
func New() *stringBuilder {
	return new(stringBuilder)
}

// TODO: how much len() cost ?
func (sb *stringBuilder) Append(s string) (int, bool) {
	if maxlength-sb.offset < len(s) {
		return 0, false
	}
	tmp := []rune(s)
	for i, v := range tmp {
		sb.buffer[sb.offset+i] = v
	}
	sb.offset += len(s)
	return len(s), true
}

func (sb *stringBuilder) Println() {
	fmt.Println(sb.ToString())
}

// TODO: what is difference bewteen append and s +=
func (sb *stringBuilder) ToString() string {
	s := ""
	for _, v := range sb.buffer {
		if v == 0 {
			break
		}
		s += string(v)
	}
	return s
}
