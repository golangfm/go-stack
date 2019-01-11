package channel

import (
	"testing"
	"time"
)

func TestChannelSelect(t *testing.T) {
	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				t.Log("ssss")
			}
		}
	}()
	time.Sleep(10 * time.Second)
	t.Fail()
}

func TestChannelRange(t *testing.T) {
	go func() {
		for _ = range time.Tick(1 * time.Second) {
			t.Log("sssss")
		}
	}()
	time.Sleep(10 * time.Second)
	t.Fail()
}

func TestRangeMap(t *testing.T) {
	data := make(map[string]string)
	data["key1"] = "val1"
	data["key2"] = "val2"
	for k, v := range data {
		t.Log(k, "\t", v)
	}
	for kk := range data {
		t.Log(kk)
	}
	t.Fail()
}
