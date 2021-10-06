package bot

import "github.com/bwmarrin/discordgo"

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func Ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "POSITIVE ATTITUDE POSITIVE MENTAL SCALING ACTUALITY")
}
