package utilities

import (
	"strings"
	"time"
)

func CheckingFileType(name string) (string, string) {
	file := strings.Split(name, ".")
	if file[len(file)-1] == "jpg" || file[len(file)-1] == "jpeg" || file[len(file)-1] == "png" {
		return file[len(file)-1], time.Now().Format("20060102150405") + "." + file[len(file)-1]
	}
	return "", ""
}
