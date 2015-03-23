package gothingspeak

import (
	"log"
	"testing"
)

func Test1(t *testing.T) {
	ts := NewChannelWriter("BZR8XQJMO65NSRTX", true)
	if !ts.AddField(1, "1000") {
		return
	}
	resp, err := ts.Update()
	if err != nil {
		log.Println("Couldn't update.")
		return
	} else {
		log.Println(resp)
	}
}
