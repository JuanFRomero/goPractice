package utils

import "net/url"

func GenerarUrl(uri, host, protocolo string, mapa map[string]string) string {
	u, _ := url.Parse(uri)
	u.Host = host
	u.Scheme = protocolo
	mapaFuncion := u.Query()

	for key, value := range mapa {
		mapaFuncion.Add(key, value)
	}
	u.RawQuery = mapaFuncion.Encode()
	return u.String()
}
