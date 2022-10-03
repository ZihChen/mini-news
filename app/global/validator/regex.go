package validator

import "regexp"

func ValidateRegex(str, rule string) (bool, error) {
	ok, err := regexp.MatchString(rule, str)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	return true, nil
}
