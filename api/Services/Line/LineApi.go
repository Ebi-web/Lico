package Line

var (
	baseEndpoint        = "https://api.line.me/v2/bot/"
	replyEndpointSuffix = "message/reply/"

	ReplyEndpoint      = baseEndpoint + replyEndpointSuffix
	ChannelAccessToken = ""
)
