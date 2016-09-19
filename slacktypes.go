package tension

// BotInfo contains the details for a particular bot.
type BotInfo struct {
	Id      UserId
	Deleted bool
	Name    string
	Icons   map[string]string
}

// ChannelMinimal is returned in a ChannelRenameResult.
type ChannelMinimal struct {
	Id         ChannelId
	Name       string
	Created    UnixTime
	Is_Channel bool
}

// ChannelBrief is returned in a ChannelJoinResult if already in a channel.
type ChannelBrief struct {
	ChannelMinimal
	Creator     UserId
	Is_Archived bool
	Is_General  bool
}

// ChannelSummary expands on ChannelBrief to include fields used by
// eg ChannelListResult.
type ChannelSummary struct {
	ChannelBrief
	Is_Member   bool
	Num_Members uint
	Topic       ChannelTopic
	Purpose     ChannelPurpose
}

// ChannelFull expands on ChannelSummary to include fields used by
// eg ChannelInfoResult.
type ChannelFull struct {
	ChannelSummary
	Is_Starred           bool
	Members              []UserId
	Last_Read            SlackTime
	Latest               Message
	Unread_Count         uint
	Unread_Count_Display uint
}

type GroupFull struct {
	Id          ChannelId
	Name        string
	Is_Group    bool
	Created     UnixTime
	Creator     UserId
	Is_Archived bool

	// Is_MPIM indicates this is a simulated private channel for an IM among 3+ users
	// This should never be true, because this only happens when API requests are made that
	// don't signal MPIM awareness, and this code is MPIM aware.
	Is_MPIM              bool
	Members              []UserId
	Topic                ChannelTopic
	Purpose              ChannelPurpose
	Last_Read            SlackTime
	Latest               Message
	Unread_Count         uint
	Unread_Count_Display uint
}

type MPIMFull struct {
	Id                   ChannelId
	Name                 string
	Is_MPIM              bool
	Is_Group             bool
	Created              UnixTime
	Creator              UserId
	Members              []UserId
	Last_Read            SlackTime
	Latest               Message
	Unread_Count         uint
	Unread_Count_Display uint
}

type IMFull struct {
	Id              ChannelId
	Is_IM           bool
	Created         UnixTime
	OtherUser       UserId `json:"user"`
	Has_Pins        bool
	Is_User_Deleted bool
	Is_Open         bool
	Last_Read       SlackTime
	Latest          SlackTime // ?!
	Unread_Count    uint
}

// ChannelTopic contains the "topic" text, as well as the user who set the topic and when
type ChannelTopic struct {
	Value    string
	Creator  UserId
	Last_Set UnixTime
}

// ChannelPurpose contains the "purpose" text, as well as the user who set the purpose and when
type ChannelPurpose struct {
	Value    string
	Creator  UserId
	Last_Set UnixTime
}

// Message describes an event in a channel or private message.
type Message struct {
	Type       string
	Timestamp  SlackTime `json:"ts"`
	User       UserId
	Text       string
	Is_Starred bool
	Reactions  []Reaction
}

// Reaction describes an emoji attachment to a message.
type Reaction struct {
	Name  string
	Count uint
	Users []UserId
}

type UserBrief struct {
	Id              UserId
	Name            string
	Prefs           interface{}
	Created         UnixTime
	Manual_Presence string
}
type UserSummary struct {
	UserBrief
	Deleted             bool
	Color               string
	Profile             map[string]string
	Is_Admin            bool
	Is_Owner            bool
	Is_Primary_Owner    bool
	Is_Restricted       bool
	Is_Ultra_Restricted bool
	Has_2FA             bool
	Two_Factor_Type     string // only present if Has_2FA is true
	Has_Files           bool
}

type TeamSummary struct {
	Id                      string
	Name                    string
	Email_Domain            string // ie anybody@(email_domain) can sign up
	Domain                  string // ie (domain).slack.com
	Icon                    IconInfo
	Msg_Edit_Window_Mins    int
	Over_Storage_Limit      bool
	Over_Integrations_Limit bool
	Plan                    string
	Prefs                   TeamPrefs
}

type TeamPrefs struct {
	Display_Email_Addresses             bool
	Allow_Shared_Channel_Perms_Override bool
	Default_Channels                    []ChannelId
	Who_Can_Create_Groups               string // "ra", ... set of runes describing capable users?
	Warn_Before_At_Channel              string   // "always", ...
	Auth_Mode                           string   // "normal", ...
	Invites_Only_Admins                 bool
	Retention_Type                      int // 0, ...
	Allow_Retention_Override            bool
	Who_Has_Team_Visibility             string // "ra", ...
	Who_Can_Kick_Channels               string // "admin", ...
	Who_Can_Edit_User_Groups            string // "admin", ...
	Msg_Edit_Window_Mins                int
	Group_Retention_Duration            int
	Group_Retention_Type                int    // 0, ...
	Allow_Shared_Channels               bool
	Invites_Limit                       bool
	Who_Can_At_Everyone                 string // "regular", ...
	Who_Can_Create_Channels             string // "regular", ...
	Who_Can_Archive_Channsl             string // "regular", ...
	Retention_Duration                  int
	Require_At_For_Mention              bool
	Who_Can_Create_Shared_Channels      string // "admin", ...
	Hide_Referers                       bool
	Calling_App_Name                    string // "Slack", ...
	Who_Can_Post_General                string // "ra", ...
	Who_Can_Kick_Groups                 string // "regular", ...
	Compliance_Export_Start             int    // 0, ...
	Allow_Message_Deletion              bool
	Allow_Calls                         bool
	Display_Real_Names                  bool
	Who_Can_At_Channel                  string // "ra", ...
	Who_Can_Create_Delete_User_Groups   string // "admin", ...
	Who_Can_Change_Team_Profile         string // "admin", ...
	
	DM_Retention_Duration               int
	DM_Retention_Type                   int    // 0, ...
	
	Disable_File_Editing                bool
	Disable_File_Deleting               bool
	Disable_File_Uploads                string // "allow_all", ...
	Disallow_Public_File_URLs           bool
	File_Retention_Duration             int
	File_Retention_Type                 int // 0, ...
	
	Default_Rxns                        []string // Rxns -> Reactions
	Team_Handy_Rxns                     HandyReactionsInfo
	Channel_Handy_Rxns                  HandyReactionsInfo
	
	DND_Start_Hour                      string // "22:00", ...
	DND_End_Hour                        string // "08:00", ...
	DND_Enabled                         bool
	
	// These are not well-documented and look mostly useless
	Who_Can_Manage_Shared_Channels      interface{}
	Who_Can_Manage_Integrations         interface{}
	Who_Can_Post_In_Shared_Channels     interface{}
}

type HandyReactionsInfo struct {
	Restrict bool
	List     []ReactionListItem
}

type ReactionListItem struct {
	Name  string
	Title string
}

type IconInfo struct {
	Image_Default bool
	Image_34      string
	Image_44      string
	Image_68      string
	Image_88      string
	Image_102     string
	Image_132     string
	Image_230     string
}

// A UserId is an opque string that represents a user or a bot.
type UserId string

// A ChannelId is an opaque string that represents a "public" Slack
// channel, survivng channel renames unchanged.
type ChannelId string

// SlackTime is a string that represents an unambiguously-sortable
// timeline event. It appears to be encoded as a unix timestamp, plus a
// dot, plus a zero-padded atomic index for each event that happens in
// that second. SlackTimes are mostly used for finding relative events
// (ie happened before, happened after) and not for precise time.
type SlackTime string

// UnixTime is simply your bog-standard Unix epoch counter, such as used
// by time.Unix. No nanoseconds.
type UnixTime uint64
