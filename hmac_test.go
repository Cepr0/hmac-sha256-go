package main

import (
	"testing"
	"time"
	"fmt"
)

func TestComputeHmac(t *testing.T) {
	actual := computeHmac("message", "secret");
	expected := "8b5f48702995c1598c573db1e21866a9b825d4a794d169d7060a03605796360b";
	if actual != expected {
		t.Errorf("Test faild!\nExpected: %s\nActual:%s\n", expected, actual)
	}
}

type TestSet struct {
	message string
	secret  string
	hmac    string
}

var test = TestSet{
	"https://www.bitmex.com/api/v1/trade/bucketed?binSize=5m&partial=false&symbol=XBTUSD&count=20&reverse=false&startTime=2018-06-02T16%3A20%3A00.000Z'",
	"8b5f48702995c1598c573db1e21866a9b825d4a794d169d7060a03605796360b",
	"2d3f7ad35a13aedfbe9601c9248761b0204120761ab3b4a2d9a6111735e272c8"}

func TestComputeHmacInCycle(t *testing.T) {
	for i := 0; i < 40; i++ {
		actual := computeHmac(test.message, test.secret);
		fmt.Println(i, ":", actual)
		if actual != test.hmac {
			t.Errorf("Test faild!\nExpected: %s\nActual:%s\n", test.hmac, actual)
		}
		time.Sleep(time.Second)
	}
}
