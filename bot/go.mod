module bot

go 1.17

require (
	github.com/bwmarrin/discordgo v0.23.2
	github.com/go-sql-driver/mysql v1.6.0
	robin/config v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/websocket v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20181030102418-4d3f4d9ffa16 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
)

replace robin/config => ../config
