package bark

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Send(host, key, title, body, group string) {
	base_url := fmt.Sprintf("https://%s/%s/%s", strings.Trim(host, "/"), key, title)

	if body != "" {
		base_url = fmt.Sprintf("%s/%s", strings.TrimSuffix(base_url, "/"), url.QueryEscape(body))
	}

	p := url.Values{}
	if group != "" {
		p.Add("group", group)
	}

	base_url = fmt.Sprintf("%s?%s", base_url, p.Encode())
	log.Println(base_url)

	resp, err := http.Get(base_url)
	if err != nil {
		log.Printf("Barksend err: %s", err)
	} else {
		log.Printf("StatusCode: %d", resp.StatusCode)
	}
}
