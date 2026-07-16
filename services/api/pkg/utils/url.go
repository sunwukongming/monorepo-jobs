package utils

import "net/url"

// URLParseQuery 获取urlquery
func URLParseQuery(query string) (map[string]string, error) {
	data, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for k, v := range data {
		m[k] = ""
		if len(v) > 0 {
			m[k] = v[0]
		}
	}
	return m, nil
}
