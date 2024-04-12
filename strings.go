package common

import s "strings"

func CombineString(strings ...string) (*string, error) {
	var builder s.Builder
	for i := range strings {
		_, err := builder.WriteString(strings[i])
		if err != nil {
			return nil, err
		}
	}
	result := builder.String()
	return &result, nil
}

func CombineStringWithSeparator(c rune, strings ...string) (*string, error) {
	var builder s.Builder
	for i := range strings {
		_, err := builder.WriteString(strings[i])
		if err != nil {
			return nil, err
		}
		_, err = builder.WriteRune(c)
		if err != nil {
			return nil, err
		}
	}
	result := builder.String()
	return &result, nil
}
