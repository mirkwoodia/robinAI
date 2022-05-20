module robin

go 1.17

require (
	robin/bot v0.0.0-00010101000000-000000000000
	robin/config v0.0.0-00010101000000-000000000000
	github.com/bwmarrin/discordgo v0.23.2 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20181030102418-4d3f4d9ffa16 // indirect
)

replace robin/bot => ./bot

replace robin/config => ./config
