package tension

import "fmt"

// BotsInfo requests information from Slack about bot named by `botid`. Use
// an empty string (ie "") to elide this optional parameter.
func (s *Slack) BotsInfo(botid string) (res *BotsInfoResult, err error) {
	res = &BotsInfoResult{}

	p := map[string]string{}
	if botid != "" {
		p["bot"] = botid
	}

	err = s.call("bots.info", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}
