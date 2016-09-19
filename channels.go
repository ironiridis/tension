package tension

import "fmt"

// ChannelList requests the list of "public" channels available on the server.
func (s *Slack) ChannelList(excludeArchived bool) (res *ChannelListResult, err error) {
	res = &ChannelListResult{}
	p := map[string]string{}

	if excludeArchived {
		p["exclude_archived"] = "1"
	}

	err = s.call("channels.list", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelInfo requests more detailed info about the channel named by `channel`.
func (s *Slack) ChannelInfo(channel ChannelId) (res *ChannelInfoResult, err error) {
	res = &ChannelInfoResult{}
	p := map[string]string{"channel": string(channel)}

	err = s.call("channels.info", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelArchive causes the public channel named by `channel` to be archived.
func (s *Slack) ChannelArchive(channel ChannelId) (res *ChannelArchiveResult, err error) {
	res = &ChannelArchiveResult{}
	p := map[string]string{"channel": string(channel)}

	err = s.call("channels.archive", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelCreate creates a new "public" channel named `name`.
func (s *Slack) ChannelCreate(name string) (res *ChannelCreateResult, err error) {
	res = &ChannelCreateResult{}
	p := map[string]string{"name": name}

	err = s.call("channels.create", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelHistory requests that the server send us events from a "public" channel
// identified by `channel`, requesting up to `count` events, optionally bounded by `oldest`
// and/or `latest` (which can be "" to exclude). Regarding the API, note that the range
// of timestamps requested will always be inclusive (we always set inclusive=1) and the
// optional unread_count_display field will always be included (we always set unreads=1).
// `count` may range from 1 to 1000; the API specifies a default of 100.
func (s *Slack) ChannelHistory(channel ChannelId, count uint, oldest, latest SlackTime) (res *ChannelHistoryResult, err error) {
	res = &ChannelHistoryResult{}
	p := map[string]string{"channel": string(channel), "inclusive": "1", "unreads": "1"}

	if oldest != "" {
		p["oldest"] = string(oldest)
	}
	if latest != "" {
		p["latest"] = string(latest)
	}
	err = s.call("channels.history", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelInvite invites a user into a "public" channel we currently occupy.
func (s *Slack) ChannelInvite(channel ChannelId, user UserId) (res *ChannelInviteResult, err error) {
	res = &ChannelInviteResult{}
	p := map[string]string{"channel": string(channel), "user": string(user)}

	err = s.call("channels.invite", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelJoin joins a new "public" channel named `name`, or creates it.
func (s *Slack) ChannelJoin(name string) (res *ChannelJoinResult, err error) {
	res = &ChannelJoinResult{}
	p := map[string]string{"name": name}

	err = s.call("channels.join", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelKick removes a user from a "public" channel.
func (s *Slack) ChannelKick(channel ChannelId, user UserId) (res *ChannelKickResult, err error) {
	res = &ChannelKickResult{}
	p := map[string]string{"channel": string(channel), "user": string(user)}

	err = s.call("channels.kick", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelLeave departs the channel named by `channel`.
func (s *Slack) ChannelLeave(channel ChannelId) (res *ChannelLeaveResult, err error) {
	res = &ChannelLeaveResult{}
	p := map[string]string{"channel": string(channel)}

	err = s.call("channels.leave", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelMark updates server as to the latest message received in the channel named by `channel`.
func (s *Slack) ChannelMark(channel ChannelId, ts SlackTime) (res *ChannelMarkResult, err error) {
	res = &ChannelMarkResult{}
	p := map[string]string{"channel": string(channel), "ts": string(ts)}

	err = s.call("channels.mark", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelRename changes the name of the channel named by `channel`.
func (s *Slack) ChannelRename(channel ChannelId, name string) (res *ChannelRenameResult, err error) {
	res = &ChannelRenameResult{}
	p := map[string]string{"channel": string(channel), "name": name}

	err = s.call("channels.rename", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelSetPurpose changes the "purpose" of the channel named by `channel`.
// The purpose may be up to 250 characters.
func (s *Slack) ChannelSetPurpose(channel ChannelId, purpose string) (res *ChannelSetPurposeResult, err error) {
	res = &ChannelSetPurposeResult{}
	p := map[string]string{"channel": string(channel), "purpose": purpose}

	err = s.call("channels.setPurpose", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelSetTopic changes the "topic" of the channel named by `channel`.
// The topic may be up to 100 characters.
func (s *Slack) ChannelSetTopic(channel ChannelId, topic string) (res *ChannelSetTopicResult, err error) {
	res = &ChannelSetTopicResult{}
	p := map[string]string{"channel": string(channel), "topic": topic}

	err = s.call("channels.setTopic", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// ChannelUnarchive causes the public channel named by `channel` to be unarchived.
func (s *Slack) ChannelUnarchive(channel ChannelId) (res *ChannelUnarchiveResult, err error) {
	res = &ChannelUnarchiveResult{}
	p := map[string]string{"channel": string(channel)}

	err = s.call("channels.unarchive", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}
