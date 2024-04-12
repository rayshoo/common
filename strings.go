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

func CombineStringReferrer(strings *[]string) (*string, error) {
	var builder s.Builder
	for i := range *strings {
		_, err := builder.WriteString((*strings)[i])
		if err != nil {
			return nil, err
		}
	}
	result := builder.String()
	return &result, nil
}

func CombineStringWithRuneSeparator(sep rune, strings ...string) (*string, error) {
	var builder s.Builder
	for i := range strings {
		_, err := builder.WriteString(strings[i])
		if err != nil {
			return nil, err
		}
		_, err = builder.WriteRune(sep)
		if err != nil {
			return nil, err
		}
	}
	result := builder.String()
	return &result, nil
}

func CombineStringWithStringSeparator(sep string, strings ...string) (*string, error) {
	var builder s.Builder
	for i := range strings {
		_, err := builder.WriteString(strings[i])
		if err != nil {
			return nil, err
		}
		_, err = builder.WriteString(sep)
		if err != nil {
			return nil, err
		}
	}
	result := builder.String()
	return &result, nil
}

func CombineStringWithRuneSeparatorReferrer(sep *rune, strings *[]string) (*string, error) {
	var builder s.Builder
	for i := range *strings {
		_, err := builder.WriteString((*strings)[i])
		if err != nil {
			return nil, err
		}
		_, err = builder.WriteRune(*sep)
		if err != nil {
			return nil, err
		}
	}
	result := builder.String()
	return &result, nil
}
func CombineStringWithStringSeparatorReferrer(sep *string, strings *[]string) (*string, error) {
	var builder s.Builder
	for i := range *strings {
		_, err := builder.WriteString((*strings)[i])
		if err != nil {
			return nil, err
		}
		_, err = builder.WriteString(*sep)
		if err != nil {
			return nil, err
		}
	}
	result := builder.String()
	return &result, nil
}
