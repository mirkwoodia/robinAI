package bot

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"
	"robin/config"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID             string
	dg                *discordgo.Session
	Db                *sql.DB
	Emoji_insert      *sql.Stmt
	EmojiStats_insert *sql.Stmt
	IsAlpha           func(s string) bool
	ContainsLiterals  *regexp.Regexp
)

// TODO: do I need to defer the dg close? Is it harming my program?
func init() {
	chars := []string{"'", "\""}
	r := strings.Join(chars, "")
	ContainsLiterals = regexp.MustCompile("[" + r + "]")
	Db = Open()
}

type Emojistruct struct {
	emoji_ID   string
	emoji_name string
	DOB        string // is this an int in sql or some other format. maybe string
}

func Start() {
	dg, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	// Register ready as a callback for the ready events.
	dg.AddHandler(Ready)

	// Register guildCreate as a callback for the guildCreate events.
        dg.AddHandler(GuildCreate)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(MessageHandler)

	// Registers guildEmojisUpdate as a callback for guildEmojiUpdate events
	dg.AddHandler(GuildEmojisUpdate)

	// Registers guildmemberadd as a callback for guildmemberadd events
	dg.AddHandler(GuildMemberAdd)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildEmojis | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Printf(`Now running. Press CTRL-C to exit.`)
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
