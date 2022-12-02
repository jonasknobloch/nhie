package application

import (
	"github.com/jonasknobloch/nhie/internal/translate"
	"golang.org/x/text/language"
	"net/http"
)

func NegotiateLanguage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Query().Has("language") {
			next.ServeHTTP(w, r)
		}

		tags, err := translate.EvaluateAcceptLanguageHeader(r.Header.Get("Accept-Language"))

		var tag language.Tag

		if err != nil {
			tag = translate.SourceLanguage
		} else {
			tag = translate.MatchLanguage(tags)
		}

		u := *r.URL
		q := u.Query()

		q.Set("language", tag.String())
		u.RawQuery = q.Encode()

		http.Redirect(w, r, u.String(), http.StatusSeeOther)
	})
}
