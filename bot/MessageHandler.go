// Make rename and checkforemoji functions to concatenate the larger func

package bot

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
	"strconv"
	"github.com/bwmarrin/discordgo"
	"robin/config"
)

// Relevant pos: 3 admin perms, 26 nickname perms,
func hasBit(n int64, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

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
		perms, e := s.State.MessagePermissions(m.Message)
		ErrorCheck(e)
		perms2 := strconv.FormatInt(perms, 2)

		if m.Content == config.BotPrefix+"permsCheck" {
			s.ChannelMessageSend(m.ChannelID, "Your perms are: "+perms2)
		}

		if strings.HasPrefix(m.Content, config.BotPrefix+"steal ") && (hasBit(perms, 3) || hasBit(perms, 30)) {
			_message := strings.Split(m.Content, config.BotPrefix+"steal ")
			if _message[1] != "" {
				name_url := strings.Split(_message[1], " ")
				// TODO: check name_url[0] for valid name, and name_url[1] for a valid URL
				if len(name_url) < 2 {
					s.ChannelMessageSend(m.ChannelID, "Failed, please check input format of r.steal name_of_emoji url_to_emoji")
				} else {
					newEmoji, e := s.GuildEmojiCreate(m.GuildID, name_url[0], imgURLtobase64(name_url[1]), nil)
					if e == nil {
						s.ChannelMessageSend(m.ChannelID, "Successfully uploaded: "+newEmoji.Name)
					} else {
						s.ChannelMessageSend(m.ChannelID, "Failed!")
					}
				}
			}
		}

		if strings.HasPrefix(m.Content, config.BotPrefix+"unsteal ") && (hasBit(perms, 3) || hasBit(perms, 30)) {
			emoji := strings.Split(m.Content, config.BotPrefix+"unsteal ")
			e = s.GuildEmojiDelete(m.GuildID, emoji[1])
			if e == nil {
				s.ChannelMessageSend(m.ChannelID, "Successfully deleted the following emoji: "+emoji[1])
			} else {
				s.ChannelMessageSend(m.ChannelID, "Command Failed")
			}
		}
		if strings.HasPrefix(m.Content, config.BotPrefix+"filtername ") && (hasBit(perms, 3)) {
			name := strings.Split(m.Content, config.BotPrefix+"filtername ")
			addNameFilter, _ := Db.Prepare("INSERT IGNORE into nameFilter_" + m.GuildID + " (name) VALUES (?);")
			_, e = addNameFilter.Exec(name[1])
			if e == nil {
				s.ChannelMessageSend(m.ChannelID, "Added " + name[1] + " to the list of filtered join names")
			}
		}
                if strings.HasPrefix(m.Content, config.BotPrefix+"unfiltername ") && (hasBit(perms, 3)) {
                        name := strings.Split(m.Content, config.BotPrefix+"unfiltername ")
                        delNameFilter, _ := Db.Prepare("DELETE IGNORE from nameFilter_" + m.GuildID + " WHERE name=(?);")
                        _, e = delNameFilter.Exec(name[1])
                        if e == nil {
                                s.ChannelMessageSend(m.ChannelID, "Deleted " + name[1] + " from the list of filtered join names")
                        } else {
				s.ChannelMessageSend(m.ChannelID, "Name is not filtered, or another error has occurred.")
			}
                }
                if strings.HasPrefix(m.Content, config.BotPrefix+"listfilterednames") && (hasBit(perms, 3)) {
                        row_names, e := Db.Query("select name from nameFilter_" + m.GuildID)
                        ErrorCheck(e)
			list_of_names := ""
			for row_names.Next() {
				temp := ""
				e = row_names.Scan(&temp)
				ErrorCheck(e)
				list_of_names += temp + "\n"
			}
			s.ChannelMessageSend(m.ChannelID, list_of_names)
                }
	}

	// check for emoji, insert into emojiDB if local emoji found
	re := regexp.MustCompile("<a?:(\\w+):(\\d{18})>")
	parts := re.FindAllStringSubmatch(m.Content, -1)
	if parts != nil {
		// parts[i][1] contains the emoji name. parts[i][2] contains the emoji ID
		Emoji_insert, _ = Db.Prepare("INSERT INTO emojis_" + m.GuildID + " (emoji_ID, emoji_name, DOB) VALUES (?, ?, ?)")
		for i := 0; i < len(parts); i++ {
			date := time.Now().Format("2006-01-02")
			_, err := s.State.Emoji(m.GuildID, parts[i][2])
			if err == nil {
				emoji_name := ContainsLiterals.ReplaceAllString(parts[i][1], "")
				println (parts[i][2], emoji_name, date)
				_, e := Emoji_insert.Exec(parts[i][2], emoji_name, date)
				ErrorCheck(e)
			}
		}
	}
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func imgURLtobase64(url string) string {
	resp, e := http.Get(url)
	ErrorCheck(e)
	defer resp.Body.Close()
	bytes, e := ioutil.ReadAll(resp.Body)
	ErrorCheck(e)
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)
	return base64Encoding
}

// Runs when the help command is called on discord (under messageHandler)
func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Just visit robinai.xyz for emoji stats (tbh dont, its poorly set up rn), and you may run r.filtername name_example to autoban anyone that joins with that name. Or r.unfiltername name. Or r.listfilterednames.")
}
