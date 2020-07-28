package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	hh, _ := strconv.ParseInt(s[:2], 10, 0)
	mm, _ := strconv.ParseInt(s[3:5], 10, 0)
	ss, _ := strconv.ParseInt(s[6:8], 10, 0)

	if strings.HasSuffix(s, "AM") {
		if hh == 12 {
			hh = 0
		}
	} else {
		if 1 <= hh && hh <= 11 {
			hh += 12
		}
	}

	return fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
}
