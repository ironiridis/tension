package tension

// APIResult is the common field set for all API result types.
type APIResult struct {
	Ok       bool   `json:"ok"`
	ErrorMsg string `json:"error"`
}

// APITestResult contains a delightful affirmation of your competence.
type APITestResult struct {
	APIResult
}

// AuthTestResult includes the response to an AuthTest API call.
type AuthTestResult struct {
	APIResult
	URL     string
	Team    string
	Name    string `json:"user"`
	Team_Id string
	User_Id UserId
}

// AuthRevokeResult includes the response to an AuthRevoke API call. Note that
// `Revoked` is useful to distinguish whether the API was called in "test mode"
type AuthRevokeResult struct {
	APIResult
	Revoked bool
}

// BotsInfoResult contains the response to a BotsInfo request.
type BotsInfoResult struct {
	APIResult
	Bot BotInfo
}

// ChannelListResult contains the response to a ChannelList request.
type ChannelListResult struct {
	APIResult
	Channels []ChannelSummary
}

// ChannelInfoResult contains the response to a ChannelInfo request.
type ChannelInfoResult struct {
	APIResult
	Channel ChannelFull
}

// ChannelArchiveResult contains the response to a ChannelArchive request.
type ChannelArchiveResult struct {
	APIResult
}

// ChannelCreateResult contains the response to a ChannelCreate request.
type ChannelCreateResult struct {
	APIResult
	Channel ChannelFull
}

// ChannelHistoryResult contains the response to a ChannelHistory request.
type ChannelHistoryResult struct {
	APIResult
	Latest     SlackTime
	Messages   []Message
	Has_More   bool
	Is_Limited bool
}

// ChannelInviteResult contains the response to a ChannelInvite request.
type ChannelInviteResult struct {
	APIResult
	Channel ChannelFull
}

// ChannelJoinResult contains the response to a ChannelJoin request.
type ChannelJoinResult struct {
	APIResult
	Already_In_Channel bool
	Channel            ChannelFull
}

// ChannelKickResult contains the response to a ChannelKick request.
type ChannelKickResult struct {
	APIResult
}

// ChannelLeaveResult contains the response to a ChannelLeave request.
type ChannelLeaveResult struct {
	APIResult
	Not_In_Channel bool
}

// ChannelMarkResult contains the response to a ChannelMark request.
type ChannelMarkResult struct {
	APIResult
}

// ChannelRenameResult contains the response to a ChannelRename request.
type ChannelRenameResult struct {
	APIResult
	Channel ChannelMinimal
}

// ChannelSetPurposeResult contains the response to a ChannelSetPurpose request.
type ChannelSetPurposeResult struct {
	APIResult
	Purpose string
}

// ChannelSetTopicResult contains the response to a ChannelSetTopic request.
type ChannelSetTopicResult struct {
	APIResult
	Topic string
}

// ChannelUnarchiveResult contains the response to a ChannelUnarchive request.
type ChannelUnarchiveResult struct {
	APIResult
}

// RTMStartResult contains the response to a RTMStart request, and a bunch of
// details about the Slack you're about to start receiving streaming events for.
type RTMStartResult struct {
	APIResult
	URL      string
	Self     UserBrief
	Team     TeamSummary
	Users    []UserSummary
	Channels []ChannelSummary
	Groups   []GroupFull
	MPIMs    []MPIMFull
	IMs      []IMFull
}
