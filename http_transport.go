package atproxy

import "net/http"

func (_ Def) HTTPTransports(
	dialers Dialers,
) (transports []*http.Transport) {
	for _, dial := range dialers {
		transports = append(transports, &http.Transport{
			DialContext: dial.DialContext,
		})
	}
	return
}
