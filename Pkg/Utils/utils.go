package Utils

import (
	"math"
	"os"
	"strings"
)

func GetEnvVariable(envName string) string {
	return os.Getenv(envName)
}

func SplitText(text string, splitLimit float64) []string {
	var splitData = make([]string, 0)
	textLen := float64(len(text))
	loop := int(math.Ceil(textLen / splitLimit))

	for i := 1; i <= loop; i++ {
		if i != loop {
			splitData = append(splitData, text[:int(splitLimit)])
			text = strings.Replace(text, text[:int(splitLimit)], "", 1)
		} else {
			splitData = append(splitData, text)
		}
	}

	return splitData
}
