package suez

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Channel struct {
	OriginEndpoint string
	DestinationURL string
}

func (c *Channel) HandlerFunc(key string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := c.buildURL(r.URL, key)
		resp, err := http.Get(url)
		if err != nil {
			log.Print(err)
			return
		}
		defer resp.Body.Close()

		var body map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			log.Print(err)
			return
		}
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Print(err)
			return
		}
	})
}

func (c *Channel) buildURL(url *url.URL, key string) string {
	return fmt.Sprint(c.DestinationURL, url.String(), fmt.Sprintf("&api_key=%s", key))
}
