package translate

import (
	"errors"
	"golang.org/x/text/language"
	"sort"
	"strconv"
	"strings"
)

type choice struct {
	lang    language.Tag
	quality float64
}

var ErrInvalidAcceptLanguageHeader = errors.New("invalid header format")

func EvaluateAcceptLanguageHeader(header string) ([]language.Tag, error) {
	header = strings.Trim(header, " ")
	elements := strings.Split(header, ",")

	choices := make([]choice, len(elements))

	for i, v := range elements {
		split := strings.Split(v, ";q=")

		if len(split) == 0 || len(split) > 2 {
			return nil, ErrInvalidAcceptLanguageHeader
		}

		var lang language.Tag
		var quality float64

		var err error

		if len(split) > 0 {
			lang, err = language.Parse(split[0])

			if err != nil {
				return nil, ErrInvalidAcceptLanguageHeader
			}
		}

		if len(split) == 1 {
			quality = 1
		}

		if len(split) == 2 {
			quality, err = strconv.ParseFloat(split[1], 64)

			if err != nil {
				return nil, ErrInvalidAcceptLanguageHeader
			}
		}

		choices[i] = choice{
			lang:    lang,
			quality: quality,
		}
	}

	sort.Slice(choices, func(i, j int) bool {
		return choices[i].quality > choices[j].quality
	})

	tags := make([]language.Tag, len(choices))

	for i, c := range choices {
		tags[i] = c.lang
	}

	return tags, nil
}
