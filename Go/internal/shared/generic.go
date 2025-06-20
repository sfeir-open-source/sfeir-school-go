package shared

import (
	"os"
)

func AppendErrorNotNil(accumulator []error, elements ...error) []error {
	var results []error
	for _, v := range elements {
		if v != nil {
			results = append(results, v)
		}
	}
	return append(accumulator, results...)
}

func AppendStringNotNil(accumulator []string, elements ...string) []string {
	var results []string
	for _, v := range elements {
		if v != "" {
			results = append(results, v)
		}
	}
	return append(accumulator, results...)
}

func GetenvOr(name string, defaultV string) string {
	var ret = os.Getenv(name)
	if ret == "" {
		ret = defaultV
	}
	return ret
}
