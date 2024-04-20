package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	re := regexp.MustCompile(`(\+?1 |1|\s|\(|\)|\-|\.|\+)`)
	phoneNumber = re.ReplaceAllString(phoneNumber, "")
	if ok, err := regexp.MatchString(`^[2-9]\d{2}[2-9]\d{2}\d{4}$`, phoneNumber); err != nil || !ok {
		return "", errors.New("phoneNumber is not correct")
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	pn, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return pn[:3], nil
}

func Format(phoneNumber string) (string, error) {
	pn, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", pn[:3], pn[3:6], pn[6:]), nil
}
