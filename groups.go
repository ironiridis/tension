package tension

import "fmt"

type GroupCreateResult struct {
	APIResult
	Group ChannelFull
}

func (s *Slack) GroupCreate(name string) (res *GroupCreateResult, err error) {
	res = &GroupCreateResult{}
	p := map[string]string{"name": name}

	err = s.call("groups.create", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}
