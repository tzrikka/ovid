package slack

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli-altsrc/v3/toml"
	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"

	"github.com/tzrikka/ovid/pkg/thrippy"
)

type API struct {
	thrippy thrippy.LinkClient
}

func LinkIDFlag(configFilePath altsrc.StringSourcer) cli.Flag {
	return &cli.StringFlag{
		Name:  "thrippy-link-slack",
		Usage: "Thrippy link ID for Slack",
		Sources: cli.NewValueSourceChain(
			cli.EnvVar("THRIPPY_LINK_SLACK"),
			toml.TOML("thrippy.links.slack", configFilePath),
		),
	}
}

func RegisterActivities(cmd *cli.Command, w worker.Worker) {
	a := API{thrippy: thrippy.NewLinkClient(cmd.String("thrippy-link-slack"), cmd)}

	registerActivity(w, a.ChatDelete, ChatDeleteName)
	registerActivity(w, a.ChatGetPermalink, ChatGetPermalinkName)
	registerActivity(w, a.ChatPostEphemeral, ChatPostEphemeralName)
	registerActivity(w, a.ChatPostMessage, ChatPostMessageName)
	registerActivity(w, a.ChatUpdate, ChatUpdateName)

	registerActivity(w, a.ConversationsArchive, ConversationsArchiveName)
	registerActivity(w, a.ConversationsClose, ConversationsCloseName)
	registerActivity(w, a.ConversationsCreate, ConversationsCreateName)
	registerActivity(w, a.ConversationsHistory, ConversationsHistoryName)
	registerActivity(w, a.ConversationsInfo, ConversationsInfoName)
	registerActivity(w, a.ConversationsInvite, ConversationsInviteName)
	registerActivity(w, a.ConversationsJoin, ConversationsJoinName)
	registerActivity(w, a.ConversationsKick, ConversationsKickName)
	registerActivity(w, a.ConversationsLeave, ConversationsLeaveName)
	registerActivity(w, a.ConversationsList, ConversationsListName)
	registerActivity(w, a.ConversationsMembers, ConversationsMembersName)
	registerActivity(w, a.ConversationsOpen, ConversationsOpenName)
	registerActivity(w, a.ConversationsRename, ConversationsRenameName)
	registerActivity(w, a.ConversationsReplies, ConversationsRepliesName)
	registerActivity(w, a.ConversationsSetPurpose, ConversationsSetPurposeName)
	registerActivity(w, a.ConversationsSetTopic, ConversationsSetTopicName)
	registerActivity(w, a.ConversationsUnarchive, ConversationsUnarchiveName)
}

func registerActivity(w worker.Worker, f any, name string) {
	w.RegisterActivityWithOptions(f, activity.RegisterOptions{Name: name})
}
