package utils

import (
	"fmt"
	"regexp"
)

func RegexpCMD(parent string, end string) (res string) {
	match := regexp.MustCompile(fmt.Sprintf(`^(.*?)%s`, end))
	param := match.FindString(parent)
	return param
}
