package slack

import (
	altsrc "github.com/urfave/cli-altsrc/v3"
	"github.com/urfave/cli-altsrc/v3/toml"
	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"

	"github.com/tzrikka/ovid/internal/thrippy"
)

type API struct {
	thrippy thrippy.LinkClient
}

// LinkIDFlag defines a CLI flag for Slack's Thrippy link ID. This flag can also
// be set using an environment variable and the application's configuration file.
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

// Register exposes Temporal activities and workflows through the Ovid worker.
func Register(cmd *cli.Command, w worker.Worker) {
	a := API{thrippy: thrippy.NewLinkClient(cmd.String("thrippy-link-slack"), cmd)}

	registerActivity(w, a.ChatDeleteActivity, ChatDeleteName)
	registerActivity(w, a.ChatGetPermalinkActivity, ChatGetPermalinkName)
	registerActivity(w, a.ChatPostEphemeralActivity, ChatPostEphemeralName)
	registerActivity(w, a.ChatPostMessageActivity, ChatPostMessageName)
	registerActivity(w, a.ChatUpdateActivity, ChatUpdateName)

	registerActivity(w, a.ConversationsArchiveActivity, ConversationsArchiveName)
	registerActivity(w, a.ConversationsCloseActivity, ConversationsCloseName)
	registerActivity(w, a.ConversationsCreateActivity, ConversationsCreateName)
	registerActivity(w, a.ConversationsHistoryActivity, ConversationsHistoryName)
	registerActivity(w, a.ConversationsInfoActivity, ConversationsInfoName)
	registerActivity(w, a.ConversationsInviteActivity, ConversationsInviteName)
	registerActivity(w, a.ConversationsJoinActivity, ConversationsJoinName)
	registerActivity(w, a.ConversationsKickActivity, ConversationsKickName)
	registerActivity(w, a.ConversationsLeaveActivity, ConversationsLeaveName)
	registerActivity(w, a.ConversationsListActivity, ConversationsListName)
	registerActivity(w, a.ConversationsMembersActivity, ConversationsMembersName)
	registerActivity(w, a.ConversationsOpenActivity, ConversationsOpenName)
	registerActivity(w, a.ConversationsRenameActivity, ConversationsRenameName)
	registerActivity(w, a.ConversationsRepliesActivity, ConversationsRepliesName)
	registerActivity(w, a.ConversationsSetPurposeActivity, ConversationsSetPurposeName)
	registerActivity(w, a.ConversationsSetTopicActivity, ConversationsSetTopicName)
	registerActivity(w, a.ConversationsUnarchiveActivity, ConversationsUnarchiveName)

	registerActivity(w, a.ReactionsAddActivity, ReactionsAddName)
	registerActivity(w, a.ReactionsGetActivity, ReactionsGetName)
	registerActivity(w, a.ReactionsListActivity, ReactionsListName)
	registerActivity(w, a.ReactionsRemoveActivity, ReactionsRemoveName)

	registerActivity(w, a.UsersConversationsActivity, UsersConversationsName)
	registerActivity(w, a.UsersGetPresenceActivity, UsersGetPresenceName)
	registerActivity(w, a.UsersIdentityActivity, UsersIdentityName)
	registerActivity(w, a.UsersInfoActivity, UsersInfoName)
	registerActivity(w, a.UsersListActivity, UsersListName)
	registerActivity(w, a.UsersLookupByEmailActivity, UsersLookupByEmailName)
	registerActivity(w, a.UsersProfileGetActivity, UsersProfileGetName)
}

func registerActivity(w worker.Worker, f any, name string) {
	w.RegisterActivityWithOptions(f, activity.RegisterOptions{Name: name})
}
