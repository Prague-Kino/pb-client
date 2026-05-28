package models

import (
	"fmt"
	"strings"
)

type Expand struct {
	Fields []string
}

func ExpandOpts(fields ...string) *Expand {
	return &Expand{
		Fields: fields,
	}
}

func (e *Expand) String() string {
	if e == nil {
		return ""
	}

	if len(e.Fields) == 0 {
		return ""
	}

	return fmt.Sprintf(
		"%s",
		strings.Join(e.Fields, ","),
	)
}
