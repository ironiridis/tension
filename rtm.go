package tension

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type RTMMessage struct {
	Type     string
	Ts       SlackTime
	Channel  ChannelId `json:",omitempty"`
	User     UserId    `json:",omitempty"`
	Text     string    `json:",omitempty"`
	Presence string    `json:",omitempty"`
}

type RTMMessageChan chan *RTMMessage

type SlackRTMChannel struct {
	c   ChannelId
	rtm *SlackRTM
	Rx  RTMMessageChan
}

type SlackRTM struct {
	slack        *Slack
	Id           chan uint64
	_id          uint64
	ws           *websocket.Conn
	rx           RTMMessageChan
	ChannelsLock sync.RWMutex
	Channels     map[ChannelId]*SlackRTMChannel
}

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

func (rtm *SlackRTM) rtmrxloop() {
	defer close(rtm.rx)
	for {
		msg := new(RTMMessage)
		err := rtm.ws.ReadJSON(msg)
		if err != nil {
			log.Println(err)
			return
		}
		rtm.rx <- msg
	}
}

func (rsr *RTMStartResult) Dial() (rtm *SlackRTM, err error) {
	rtm = &SlackRTM{rx: make(RTMMessageChan), Id: make(chan uint64)}
	d := &websocket.Dialer{}
	rtm.ws, _, err = d.Dial(rsr.URL, nil)
	if err != nil {
		return
	}

	go rtm.rtmrxloop()
	go func() {
		for {
			rtm._id++
			rtm.Id <- rtm._id
		}
	}()
	return
}

type RTMSendMessageRequest struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (rtm *SlackRTM) OpenChannel(c ChannelId) (rtmc *SlackRTMChannel) {
	rtmc = &SlackRTMChannel{
		Rx:  make(RTMMessageChan),
		rtm: rtm,
		c:   c,
	}
	rtm.ChannelsLock.Lock()
	rtm.Channels[c] = rtmc
	rtm.ChannelsLock.Unlock()
	return
}

func (rtmc *SlackRTMChannel) Close() {
	rtm := rtmc.rtm
	rtm.ChannelsLock.Lock()
	delete(rtm.Channels, rtmc.c)
	rtm.ChannelsLock.Unlock()
	close(rtmc.Rx)
	return
}

func (rtmc *SlackRTMChannel) Send(text string) (err error) {
	req := &RTMSendMessageRequest{
		Id:      <-rtmc.rtm.Id,
		Type:    "message",
		Channel: string(rtmc.c),
		Text:    text,
	}
	err = rtmc.rtm.ws.WriteJSON(req)
	return
}
