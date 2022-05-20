package bot

import (
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler in bot.go) every time a new
// guild is joined or reconnected to.
func GuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	// TODO: Check if this knows when the bot gets kicked/leaves a server
	if event.Guild.Unavailable {
	//	delete_server_from_table, e := Db.Query("DELETE IGNORE FROM servers where server_ID=" + event.ID + ";")
	//	ErrorCheck(e)
	//	delete_server_from_table.Close()
		return
	}

	// Add server to servers table
	add_to_server_table, _ := Db.Prepare("INSERT IGNORE INTO servers (server_ID, server_name) VALUES (?, ?)")
	server_name := ContainsLiterals.ReplaceAllString(event.Name, "")
	_, e := add_to_server_table.Exec(event.ID, server_name)
	ErrorCheck(e)
	add_to_server_table.Close()

	println ("guild ID: " + event.ID)


        //initialize table for emoji tracking
        create_nameFilterTable, e := Db.Query("CREATE TABLE IF NOT EXISTS nameFilter_" + event.ID + " (name VARCHAR(32));")
        ErrorCheck(e)
        create_nameFilterTable.Close()

	//initialize table for emoji tracking
	create_emojiTable, e := Db.Query("CREATE TABLE IF NOT EXISTS emojis_" + event.ID + " (emoji_ID VARCHAR(20), emoji_name VARCHAR(32), DOB DATE);")
	ErrorCheck(e)
	create_emojiTable.Close()

	// initialize emoji table to store current guild emojis, to figure out which ones are unused
	create_emojiGuild, e := Db.Query("CREATE TABLE IF NOT EXISTS emojiGuild_" + event.ID + " (emoji_ID VARCHAR(20), emoji_name VARCHAR(32));")
	ErrorCheck(e)
	create_emojiGuild.Close()

	// This part takes the guild's emojis and stores them into emojiGuild_ID database
	EmojiGuild_insert, _ := Db.Prepare("INSERT IGNORE INTO emojiGuild_" + event.ID + " (emoji_ID, emoji_name) VALUES (?, ?)")
	emojiArray, _ := s.GuildEmojis(event.ID)
	for i := 0; i < len(emojiArray); i++ {
		emoji_id := emojiArray[i].ID
		emoji_name := emojiArray[i].Name
		_, e := EmojiGuild_insert.Exec(emoji_id, emoji_name)
		ErrorCheck(e)

	}
	EmojiGuild_insert.Close()
}

// This is the handler for the guild emojis update event
func GuildEmojisUpdate(s *discordgo.Session, event *discordgo.GuildEmojisUpdate) {
	UpdateEmojis(s, event.GuildID)
}

func UpdateEmojis(s *discordgo.Session, guildID string) {
	// Truncate the emojiGuild_ table and we'll reinsert the guild emojis from scratch
	delete_emojiGuild, e := Db.Query("TRUNCATE emojiGuild_" + guildID + ";")
	ErrorCheck(e)
	defer delete_emojiGuild.Close()

	EmojiGuild_insert, _ := Db.Prepare("INSERT IGNORE INTO emojiGuild_" + guildID + " (emoji_ID, emoji_name) VALUES (?, ?)")
	emojiArray, _ := s.GuildEmojis(guildID)
	for i := 0; i < len(emojiArray); i++ {
		emoji_id := emojiArray[i].ID
		emoji_name := emojiArray[i].Name
		_, e := EmojiGuild_insert.Exec(emoji_id, emoji_name)
		ErrorCheck(e)
	}
	EmojiGuild_insert.Close()
}

func GuildMemberAdd (s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	if nameExists(event.Member.User.Username, event.Member.GuildID) {
		err := s.GuildBanCreateWithReason(event.Member.GuildID, event.Member.User.ID, "Entered server with a filtered name", 1)
		if err != nil {
			println("An error has occured: unable to ban member " + event.Member.User.ID + " triggering name filter in guildID: " + event.Member.GuildID)
		}
	}
}

func nameExists(name string, guildID string) bool {
    row := Db.QueryRow("select name from nameFilter_" + guildID + " where name= ?", name)
    temp := ""
    row.Scan(&temp)
    if temp != "" {
        return true
    }
    return false
}
