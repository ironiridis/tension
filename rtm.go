package tension

import "fmt"

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
