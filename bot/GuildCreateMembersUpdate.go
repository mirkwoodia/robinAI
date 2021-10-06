package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// checks when members update nicknames/names, to see if they should be added on or removed from nonamesDB
func MembersUpdate(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	if checkForBadName(m.Member) {
		_, e := Nonames_insert.Exec(m.User.ID)
		ErrorCheck(e)
	} else {
		nonameDelete, e := Db.Query("delete ignore from nonames_" + m.Member.GuildID + " where userids=" + m.Member.User.ID)
		ErrorCheck(e)
		nonameDelete.Close()
	}
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined or reconnected to.
func GuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	// initialize table nonamesDB
	create_nonamesTable, e := Db.Query("CREATE TABLE IF NOT EXISTS nonames_" + event.ID + " (userids BIGINT PRIMARY KEY);")
	ErrorCheck(e)
	create_nonamesTable.Close()

	//initialize table emojiDB
	create_emojiTable, e := Db.Query("CREATE TABLE IF NOT EXISTS emojis_" + event.ID + " (emoji_ID BIGINT, emoji_name VARCHAR(32), DOB DATE);")
	ErrorCheck(e)
	create_emojiTable.Close()

	// initialize table emojiCounter for use in front end
	create_emojiStats, e := Db.Query("CREATE TABLE IF NOT EXISTS emojistats_" + event.ID + " (emoji_ID BIGINT, emoji_name VARCHAR(32), 1m INT, 3m INT, 6m INT, 12m INT);")
	ErrorCheck(e)
	create_emojiStats.Close()

	// if bot restarts, clean the old db first of normal names. I guess thats like select * from nonamesdb, then check nicks, and if they dont have a nick, check name, and if it returns a nonalpha 1st char, then remove from nonamesdb
	cleanNonames(event.ID, s)

	// Now run through member list and add bad usernames to the db
	Nonames_insert, _ = Db.Prepare("INSERT IGNORE INTO nonames_" + event.Guild.ID + " (userids) VALUES (?)")
	memberList, _ := s.GuildMembers(event.Guild.ID, "", 1000)
	memListRunner(s, event.Guild.ID, memberList)
}

// Returns the table data into a nonames type
func cleanNonames(guildID string, s *discordgo.Session) {
	var nonames = Nonamestruct{}
	row_nonames, e := Db.Query("select * from nonames_" + guildID)
	ErrorCheck(e)
	for row_nonames.Next() {
		e = row_nonames.Scan(&nonames.ID)
		ErrorCheck(e)

		// If user is not in guild, this will give an error. Check it differently, so like, if error: remove from db
		member, e := s.GuildMember(guildID, strconv.Itoa(nonames.ID)) // Itoa? is that int to alpha?
		ErrorCheck(e)
		if !checkForBadName(member) {
			// he has a good name, so lets remove him for the db
			_, e := Db.Query("delete from nonames_" + guildID + " where userids=" + strconv.Itoa(nonames.ID))
			ErrorCheck(e)
		}
	}
	row_nonames.Close()
}

// Returns True if name is bad (the first letter of nick/andor name is nonalpha)
func checkForBadName(m *discordgo.Member) bool {
	if m.Nick != "" {
		if !IsAlpha(string([]rune(m.Nick)[0])) {
			return true
		}
	} else if !IsAlpha(string([]rune(m.User.Username)[0])) {
		return true
	}
	return false
}

// Runs through the member list, adding any sus names to the nonames db. Recursively run through memberList because the limit is 1000 people at a time.
func memListRunner(s *discordgo.Session, guildID string, memberList []*discordgo.Member) {

	// Now we can loop through memberList, check for nicks/names to database, then the next for loop will go through the next 1000 or it will be nil and end itself
	for i := 0; i < len(memberList); i++ {
		if checkForBadName(memberList[i]) {
			_, e := Nonames_insert.Exec(memberList[i].User.ID)
			ErrorCheck(e)
		}
	}
	// Get the lastMember from the prev memberList, and get a new memberList for the ones after lastMember
	lastMember := memberList[len(memberList)-1].User.ID
	memberList, _ = s.GuildMembers(guildID, lastMember, 1000)
	if len(memberList) > 0 {
		memListRunner(s, guildID, memberList)
	}
}
