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

func Test2(t *testing.T) {
	ts := NewChannelWriter("BZR8XQJMO65NSRTX", true)
	if ts.AddField(0, "1000") {
		t.Error("Succeeded adding filed 0")
	}

	if ts.AddField(8, "1000") {
		t.Error("Succeeded adding filed 8")
	}
}
