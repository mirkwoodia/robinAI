package bot

import (
	//"database/sql"

	"database/sql"
	"fmt"
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/mirkwoodia/RobinAI/config"
)

var (
	BotID             string
	dg                *discordgo.Session
	Db                *sql.DB
	Emoji_insert      *sql.Stmt
	Nonames_insert    *sql.Stmt
	EmojiStats_insert *sql.Stmt
	IsAlpha           func(s string) bool
)

// TODO: do I need to defer the dg close? Is it harming my program?
func init() {
	IsAlpha = regexp.MustCompile(`[a-zA-Z]`).MatchString
	Db = Open()
	// defer dg.Close()
	// Db.Close()
}

type Emojistruct struct {
	emoji_ID   int
	emoji_name string
	DOB        string // is this an int in sql or some other format. maybe string
}

type Nonamestruct struct {
	ID int
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

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(MessageHandler)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(GuildCreate)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(MembersUpdate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildEmojis | discordgo.IntentsGuildMembers | discordgo.IntentsAll

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
		panic(err.Error())
	}
}
