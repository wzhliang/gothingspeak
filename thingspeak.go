package gothingspeak

// http://community.thingspeak.com/documentation/api/

import (
	"fmt"
	"log"
	"net/http"
)

const (
	apiBase        string = "api.thingspeak.com"
	MAX_LEN_FIELDS int    = 8
)

type ChannelWriter struct {
	tls       bool
	key       string
	fields    map[int]string
	lat       string
	long      string
	elevation string
	path      string
	url       string
}

func NewChannelWriter(key string, useTls bool) *ChannelWriter {
	w := new(ChannelWriter)
	w.tls = useTls
	w.key = key
	w.fields = make(map[int]string)

	return w
}

func (w *ChannelWriter) buildUrl() {
	var schema string

	if w.tls {
		schema = "https"
	} else {
		schema = "http"
	}
	w.path = "update"

	w.url = fmt.Sprintf("%s://%s/%s?api_key=%s", schema, apiBase, w.path, w.key)
	for n, f := range w.fields {
		w.url += fmt.Sprintf("&field%d=%s", n, f)
	}
}

func (w *ChannelWriter) AddField(n int, value string) bool {
	if n <= 0 || n >= MAX_LEN_FIELDS {
		return false // FIXME: return error
	}

	w.fields[n] = value

	return true
}

func (w *ChannelWriter) Update() (resp *http.Response, err error) {
	w.buildUrl()
	log.Printf("URL: %s\n", w.url)

	client := &http.Client{}
	r, err := http.NewRequest("POST", w.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err = client.Do(r)

	return
}

// EoF
