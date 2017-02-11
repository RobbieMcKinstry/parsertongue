package main

import (
	"fmt"
)

func (t Token) String() string {
	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}
