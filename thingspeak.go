package gothingspeak

import (
	"fmt"
	"log"
	"net/http"
)

const baseUrl string = "https://api.thingspeak.com"
const updatePath string = "update"

func buildUpdateUrl(key string, field string, value string) string {
	s := fmt.Sprintf("%s/%s?api_key=%s&%s=%s", baseUrl, updatePath, key, field, value)
	log.Println(s)

	return s
}

func Update(key string, field string, value string) (resp *http.Response, err error) {
	client := &http.Client{}
	r, err := http.NewRequest("POST", buildUpdateUrl(key, field, value), nil)
	if err != nil {
		return nil, err
	}
	resp, err = client.Do(r)

	return
}

/* From thingspeak doc, updating a channel:
https://api.thingspeak.com/update?api_key=YOUR_CHANNEL_API_KEY&field1=7
*/
