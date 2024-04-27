package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

type Entry struct {
	Date        string
	Description string
	Change      int
}

type Locale struct {
	name                 string
	headerDate           string
	headerDescription    string
	headerChange         string
	separatorCurrency    string
	separatorDecimal     string
	separatorThousand    string
	dateFormat           func(string, string, string) string
	positiveNumberFormat func(string, string) string
	negativeNumberFormat func(string, string) string
}

var currencySigns map[string]string

func init() {
	currencySigns = map[string]string{
		"USD": "$",
		"EUR": "€",
	}
}

func locale(name string) *Locale {
	switch name {
	case "en-US":
		return &Locale{
			name:                 "en-US",
			headerDate:           "Date",
			headerDescription:    "Description",
			headerChange:         "Change",
			separatorCurrency:    "",
			separatorDecimal:     ".",
			separatorThousand:    ",",
			dateFormat:           func(day, month, year string) string { return fmt.Sprintf("%s/%s/%s", month, day, year) },
			positiveNumberFormat: func(number, sign string) string { return fmt.Sprintf("%s%s ", sign, number) },
			negativeNumberFormat: func(number, sign string) string { return fmt.Sprintf("(%s%s)", sign, number) },
		}
	case "nl-NL":
		return &Locale{
			name:                 "nl-NL",
			headerDate:           "Datum",
			headerDescription:    "Omschrijving",
			headerChange:         "Verandering",
			separatorCurrency:    " ",
			separatorDecimal:     ",",
			separatorThousand:    ".",
			dateFormat:           func(day, month, year string) string { return fmt.Sprintf("%s-%s-%s", day, month, year) },
			positiveNumberFormat: func(number, sign string) string { return fmt.Sprintf("%s %s ", sign, number) },
			negativeNumberFormat: func(number, sign string) string { return fmt.Sprintf("%s %s-", sign, number) },
		}
	default:
		return nil
	}
}

func compare(first, second Entry) bool {
	if first.Date != second.Date {
		return first.Date < second.Date
	}
	if first.Description != second.Description {
		return first.Description < second.Description
	}
	if first.Change != second.Change {
		return first.Change < second.Change
	}
	return false
}
func formatRow(cells []string, widths []int) string {
	result := ""
	for i, text := range cells {
		right := false
		width := widths[i]
		if width < 0 {
			width *= -1
			right = true
		}
		length := utf8.RuneCountInString(text)
		if length > width {
			text = text[0:len(text)-5] + "..."
			length = width
		}
		if i != 0 {
			result += " | "
		}
		if right {
			result += strings.Repeat(" ", width-length) + text
		} else if i != len(cells)-1 {
			result += text + strings.Repeat(" ", width-length)
		} else {
			result += text
		}
	}
	return result + "\n"
}
func formatNumber(n int, currencySign string, locale Locale) string {
	cents := n
	if cents < 0 {
		cents *= -1
	}
	number := fmt.Sprintf("%s%02d", locale.separatorDecimal, cents%100)
	cents /= 100
	number = fmt.Sprintf("%d%s", cents%1000, number)
	cents /= 1000
	for cents > 0 {
		number = fmt.Sprintf("%d%s%s", cents%1000, locale.separatorThousand, number)
		cents /= 1000
	}
	if n >= 0 {
		return locale.positiveNumberFormat(number, currencySign)
	} else {
		return locale.negativeNumberFormat(number, currencySign)
	}
}
func formatDate(locale *Locale, date string) (string, error) {
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		return "", errors.New("")
	}
	year, month, day := date[0:4], date[5:7], date[8:10]
	return locale.dateFormat(day, month, year), nil
}
func FormatLedger(currency string, localeName string, entries []Entry) (string, error) {
	locale := locale(localeName)
	currencySign, validCurrency := currencySigns[currency]
	if locale == nil || !validCurrency {
		return "", errors.New("")
	}
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.SliceStable(entriesCopy, func(i, j int) bool {
		return compare(entriesCopy[i], entriesCopy[j])
	})
	result := formatRow([]string{locale.headerDate, locale.headerDescription, locale.headerChange}, []int{10, 25, 13})
	for _, entry := range entriesCopy {
		date, err := formatDate(locale, entry.Date)
		if err != nil {
			return "", errors.New("")
		}
		amount := formatNumber(entry.Change, currencySign, *locale)
		result += formatRow([]string{date, entry.Description, amount}, []int{10, 25, -13})
	}
	return result, nil
}
