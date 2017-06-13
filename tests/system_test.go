package system

import (
	"bytes"
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

	time.Sleep(time.Second * 5)
	// create a channel to send results
	// resultChan := make(chan Score)

	// connect to the result events
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("result", 80, false),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		t.Error(err)
	}
	c.On(gosocketio.OnConnection, func(c *gosocketio.Channel, args interface{}) {
		t.Log("connected")
		c.Join("scores")
	})
	c.On("scores", func(c *gosocketio.Channel, args interface{}) {
		t.Log(args)
		//resultChan <- score
	})
	defer c.Close()

	// vote for a category
	postData := url.Values{}
	postData.Add("vote", "a")

	resp, err := http.Post("http://vote", "application/x-www-form-urlencoded", bytes.NewBufferString(postData.Encode()))
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("got non OK status", resp.StatusCode)
	}

	time.Sleep(time.Second * 10)

	t.Error("error")
	// score := <-resultChan
	// if score.A != 1 {
	// 	t.Error("score was not expected value", score.A)
	// }
}
