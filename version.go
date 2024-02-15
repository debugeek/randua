package randua

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type VersionRange struct {
	Min int
	Max int
}

func BuildVersion(separator string, params ...any) string {
	var values []string
	for _, param := range params {
		switch value := param.(type) {
		case VersionRange:
			values = append(values, fmt.Sprintf("%d", value.Min+rand.Intn(value.Max-value.Min+1)))
		case string:
			values = append(values, value)
		}
	}
	return Concat(separator, values)
}

func Concat(separator string, values ...any) string {
	var elements []string
	for _, value := range values {
		switch value := value.(type) {
		case string:
			elements = append(elements, value)
		case []string:
			elements = append(elements, value...)
		default:
			log.Print(value)
		}
	}
	return strings.Join(elements, separator)
}
