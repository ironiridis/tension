package tension

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// RTMStart request a special Websocket URL plus a ton of anciliary information
// about the Slack instance we'll be talking to.
// `simpleLatest` will elide the timestamps for all but the most recent messages in each channel.
// `noUnreads` will skip counts for unread messages per channel.
func (s *Slack) RTMStart(simpleLatest, noUnreads bool) (res *RTMStartResult, err error) {
	res = &RTMStartResult{}
	p := map[string]string{"mpim_aware": "1"}

	if simpleLatest {
		p["simple_latest"] = "1"
	}

	if noUnreads {
		p["no_unreads"] = "1"
	}

	err = s.call("rtm.start", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}

	return
}

type RTMMessage struct {
	Type    string
	Ts      SlackTime
	Channel string `json:",omitempty"`
	User    string `json:",omitempty"`
	Text    string `json:",omitempty"`
}

func (rtm *SlackRTM) rtmrxloop() {
	defer close(rtm.Rx)
	for {
		msg := new(RTMMessage)
		err := rtm.ws.ReadJSON(msg)
		if err != nil {
			log.Println(err)
			return
		}
		rtm.Rx <- msg
	}
}

func (rsr *RTMStartResult) Dial() (rtm *SlackRTM, err error) {
	rtm = &SlackRTM{Rx: make(chan *RTMMessage)}
	d := &websocket.Dialer{}
	rtm.ws, _, err = d.Dial(rsr.URL, nil)
	go rtm.rtmrxloop()
	return
}

type RTMSendMessageRequest struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (rtm *SlackRTM) SendMessage(dest, text string) (err error) {
	rtm.id++
	req := &RTMSendMessageRequest{
		Id:      rtm.id,
		Type:    "message",
		Channel: dest,
		Text:    text,
	}
	err = rtm.ws.WriteJSON(req)
	return
}
