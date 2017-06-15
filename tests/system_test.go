package system

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Score struct {
	A int
	B int
}

func TestVote(t *testing.T) {

	// Wait for the other services to startup
	time.Sleep(time.Second * 10)

	//
	//	Listen for Results
	//

	resultChan := make(chan Score)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("result", 80, false),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		t.Fatal(err)
	}
	c.On(gosocketio.OnConnection, func(c *gosocketio.Channel, args interface{}) {
		t.Log("connected")
		c.Join("scores")
	})
	c.On("scores", func(c *gosocketio.Channel, args string) {
		var score Score
		if err := json.Unmarshal([]byte(args), &score); err != nil {
			t.Fatal(err)
		}
		resultChan <- score
	})
	defer c.Close()

	//
	//	Vote for A
	//

	postData := url.Values{}
	postData.Add("vote", "a")

	resp, err := http.Post("http://vote", "application/x-www-form-urlencoded", bytes.NewBufferString(postData.Encode()))
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("got non OK status", resp.StatusCode)
	}

	//
	//	Validate result
	//

	select {
	case score := <-resultChan:
		if score.A < 1 {
			t.Error("score was not expected value", score.A)
		}
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting for result")
	}
}
