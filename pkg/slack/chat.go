package slack

import (
	"context"
	"errors"
	"net/url"
)

const (
	ChatDeleteName        = "slack.chat.delete"
	ChatGetPermalinkName  = "slack.chat.getPermalink"
	ChatPostEphemeralName = "slack.chat.postEphemeral"
	ChatPostMessageName   = "slack.chat.postMessage"
	ChatUpdateName        = "slack.chat.update"
)

// https://docs.slack.dev/reference/methods/chat.delete
func (a *API) ChatDelete(ctx context.Context, req *ChatDeleteRequest) (*ChatDeleteResponse, error) {
	resp := new(ChatDeleteResponse)
	if err := a.httpPost(ctx, ChatDeleteName, req, resp); err != nil {
		return nil, err
	}
	if !resp.OK {
		return nil, errors.New("Slack API error: " + resp.Error)
	}
	return resp, nil
}

// https://docs.slack.dev/reference/methods/chat.delete
type ChatDeleteRequest struct {
	Channel string `json:"channel"`
	TS      string `json:"ts"`

	AsUser bool `json:"as_user,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.delete
type ChatDeleteResponse struct {
	slackResponse

	Channel string `json:"channel,omitempty"`
	TS      string `json:"ts,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.getPermalink
func (a *API) ChatGetPermalink(ctx context.Context, req *ChatGetPermalinkRequest) (*ChatGetPermalinkResponse, error) {
	query := url.Values{}
	query.Set("channel", req.Channel)
	query.Set("message_ts", req.MessageTS)

	resp := new(ChatGetPermalinkResponse)
	if err := a.httpGet(ctx, ChatGetPermalinkName, query, resp); err != nil {
		return nil, err
	}
	if !resp.OK {
		return nil, errors.New("Slack API error: " + resp.Error)
	}
	return resp, nil
}

// https://docs.slack.dev/reference/methods/chat.getPermalink
type ChatGetPermalinkRequest struct {
	Channel   string `json:"channel"`
	MessageTS string `json:"message_ts"`
}

// https://docs.slack.dev/reference/methods/chat.getPermalink
type ChatGetPermalinkResponse struct {
	slackResponse

	Channel   string `json:"channel,omitempty"`
	Permalink string `json:"permalink,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postEphemeral
func (a *API) ChatPostEphemeral(ctx context.Context, req *ChatPostEphemeralRequest) (*ChatPostEphemeralResponse, error) {
	resp := new(ChatPostEphemeralResponse)
	if err := a.httpPost(ctx, ChatPostEphemeralName, req, resp); err != nil {
		return nil, err
	}
	if !resp.OK {
		return nil, errors.New("Slack API error: " + resp.Error)
	}
	return resp, nil
}

// https://docs.slack.dev/reference/methods/chat.postEphemeral
//
// https://docs.slack.dev/reference/methods/chat.postMessage#channels
type ChatPostEphemeralRequest struct {
	Channel string `json:"channel"`
	User    string `json:"user"`

	Attachments  []map[string]any `json:"attachments,omitempty"`
	Blocks       []map[string]any `json:"blocks,omitempty"`
	IconEmoji    string           `json:"icon_emoji,omitempty"`
	IconURL      string           `json:"icon_url,omitempty"`
	LinkNames    bool             `json:"link_names,omitempty"`
	MarkdownText string           `json:"markdown_text,omitempty"`
	Parse        string           `json:"parse,omitempty"`
	Text         string           `json:"text,omitempty"`
	ThreadTS     string           `json:"thread_ts,omitempty"`
	Username     string           `json:"username,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postEphemeral
type ChatPostEphemeralResponse struct {
	slackResponse

	MessageTS string `json:"message_ts,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postMessage
func (a *API) ChatPostMessage(ctx context.Context, req *ChatPostMessageRequest) (*ChatPostMessageResponse, error) {
	resp := new(ChatPostMessageResponse)
	if err := a.httpPost(ctx, ChatPostMessageName, req, resp); err != nil {
		return nil, err
	}
	if !resp.OK {
		return nil, errors.New("Slack API error: " + resp.Error)
	}
	return resp, nil
}

// https://docs.slack.dev/reference/methods/chat.postMessage
type ChatPostMessageRequest struct {
	Channel string `json:"channel"`

	Attachments  []map[string]any `json:"attachments,omitempty"`
	Blocks       []map[string]any `json:"blocks,omitempty"`
	IconEmoji    string           `json:"icon_emoji,omitempty"`
	IconURL      string           `json:"icon_url,omitempty"`
	LinkNames    bool             `json:"link_names,omitempty"`
	MarkdownText string           `json:"markdown_text,omitempty"`
	Metadata     map[string]any   `json:"metadata,omitempty"`
	// Ignoring "mrkdwn" for now, because it has an unusual default value (true).
	Parse          string `json:"parse,omitempty"`
	ReplyBroadcast bool   `json:"reply_broadcast,omitempty"`
	Text           string `json:"text,omitempty"`
	ThreadTS       string `json:"thread_ts,omitempty"`
	UnfurnLinks    bool   `json:"unfurl_links,omitempty"`
	Username       string `json:"username,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postMessage
type ChatPostMessageResponse struct {
	slackResponse

	Channel string         `json:"channel,omitempty"`
	TS      string         `json:"ts,omitempty"`
	Message map[string]any `json:"message,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.update
func (a *API) ChatUpdate(ctx context.Context, req *ChatUpdateRequest) (*ChatUpdateResponse, error) {
	resp := new(ChatUpdateResponse)
	if err := a.httpPost(ctx, ChatUpdateName, req, resp); err != nil {
		return nil, err
	}
	if !resp.OK {
		return nil, errors.New("Slack API error: " + resp.Error)
	}
	return resp, nil
}

// https://docs.slack.dev/reference/methods/chat.update
//
// https://docs.slack.dev/reference/methods/chat.postMessage#channels
type ChatUpdateRequest struct {
	Channel string `json:"channel"`
	TS      string `json:"ts"`

	Attachments    []map[string]any `json:"attachments,omitempty"`
	Blocks         []map[string]any `json:"blocks,omitempty"`
	MarkdownText   string           `json:"markdown_text,omitempty"`
	Metadata       map[string]any   `json:"metadata,omitempty"`
	LinkNames      bool             `json:"link_names,omitempty"`
	Parse          string           `json:"parse,omitempty"`
	Text           string           `json:"text,omitempty"`
	ReplyBroadcast bool             `json:"reply_broadcast,omitempty"`
	FileIDs        []string         `json:"file_ids,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.update
type ChatUpdateResponse struct {
	slackResponse

	Channel string         `json:"channel,omitempty"`
	TS      string         `json:"ts,omitempty"`
	Text    string         `json:"text,omitempty"`
	Message map[string]any `json:"message,omitempty"`
}
