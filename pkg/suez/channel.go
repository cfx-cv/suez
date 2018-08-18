package suez

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/cfx-cv/herald/pkg/common"
)

type Channel struct {
	OriginEndpoint string
	DestinationURL string

	Method string
}

func (c *Channel) HandlerFunc(envs map[string]string) http.HandlerFunc {
	switch c.Method {
	default:
		return c.handlerFuncGET(envs)
	}
}

func (c *Channel) handlerFuncGET(envs map[string]string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := c.buildURL(r.URL, envs)
		resp, err := http.Get(url)
		if err != nil {
			log.Print(err)
			common.Publish(common.SuezErrors, err.Error())
			return
		}
		defer resp.Body.Close()

		var body map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			log.Print(err)
			common.Publish(common.SuezErrors, err.Error())
			return
		}
		if err := json.NewEncoder(w).Encode(body); err != nil {
			log.Print(err)
			common.Publish(common.SuezErrors, err.Error())
			return
		}
	})
}

func (c *Channel) buildURL(url *url.URL, envs map[string]string) string {
	key := envs["GCP_API_KEY"]
	return fmt.Sprint(c.DestinationURL, url.String(), fmt.Sprintf("&api_key=%s", key))
}
