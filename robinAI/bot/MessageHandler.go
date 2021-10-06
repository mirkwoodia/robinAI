// Make rename and checkforemoji functions to concatenate the larger func

package bot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mirkwoodia/RobinAI/config"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		if m.Content == config.BotPrefix+"help" {
			helpCommand(s, m)
		}

		if m.Content == config.BotPrefix+"testdb" {
			TestTables(s, m)
		}

		if strings.HasPrefix(m.Content, config.BotPrefix+"noname ") {
			totalRows, _ := Db.Query("select count(*) from nonames_" + m.GuildID)
			var userCount = ""
			totalRows.Next()
			totalRows.Scan(&userCount)
			totalRows.Close()
			s.ChannelMessageSend(m.ChannelID, "Renaming "+userCount+" users...")
			newNick := strings.Split(m.Content, config.BotPrefix+"noname ")
			// loop through nonamesdb and
			var nonames = Nonamestruct{}
			row_nonames, e := Db.Query("select * from nonames_" + m.GuildID)
			ErrorCheck(e)
			failures := 0
			for row_nonames.Next() {
				e = row_nonames.Scan(&nonames.ID)
				ErrorCheck(e)
				e := s.GuildMemberNickname(m.GuildID, strconv.Itoa(nonames.ID), strings.Trim(newNick[1], " "))
				if e != nil {
					failures++
				}
			}
			if failures == 0 {
				s.ChannelMessageSend(m.ChannelID, "Finished renaming "+userCount+" users!")
			} else {
				users, _ := strconv.Atoi(userCount)
				successes := users - failures
				s.ChannelMessageSend(m.ChannelID, "Finished renaming "+strconv.Itoa(successes)+" users! Failures: "+strconv.Itoa(failures))
			}
			row_nonames.Close()
		}
	}

	// check for emoji, insert into emojiDB if local emoji found
	Emoji_insert, _ = Db.Prepare("INSERT INTO emojis_" + m.GuildID + " (emoji_ID, emoji_name, DOB) VALUES (?, ?, ?)")
	re := regexp.MustCompile("<a?:(\\w+):(\\d{18})>")
	parts := re.FindAllStringSubmatch(m.Content, -1)
	if parts != nil {
		for i := 0; i < len(parts); i++ {
			date := time.Now().Format("2006-01-02")
			emojiID, _ := strconv.Atoi(parts[i][2])
			_, err := s.State.Emoji(m.GuildID, parts[i][2])
			if err == nil {
				_, e := Emoji_insert.Exec(emojiID, parts[i][1], date)
				ErrorCheck(e)
			}
		}
	}
}

// Runs when the help command is called on discord (under messageHandler)
func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Help command is under construction, please wait for a future release")
}

// Runs when the testdb command is called on discord (under messageHandler)
func TestTables(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Sending test results to command line")

	row_nonames, e := Db.Query("select * from nonames_" + m.GuildID)
	ErrorCheck(e)
	row_emojis, e := Db.Query("select * from emojis_" + m.GuildID)
	ErrorCheck(e)

	var nonames = Nonamestruct{}
	var emojis = Emojistruct{}

	for row_nonames.Next() {
		e = row_nonames.Scan(&nonames.ID)
		ErrorCheck(e)
		fmt.Println(nonames)
	}

	for row_emojis.Next() {
		e = row_emojis.Scan(&emojis.emoji_ID, &emojis.emoji_name, &emojis.DOB)
		ErrorCheck(e)
		fmt.Println(emojis)
	}
	row_nonames.Close()
	row_emojis.Close()
}
