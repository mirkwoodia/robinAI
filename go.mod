module robin

go 1.15

require (
	github.com/aws/aws-sdk-go v1.40.28
	github.com/bwmarrin/discordgo v0.23.2
	github.com/go-sql-driver/mysql v1.6.0
	robin/bot v0.0.0-00010101000000-000000000000 // indirect
	robin/config v0.0.0-00010101000000-000000000000 // indirect
)

replace robin/bot => /bot

replace robin/config => /config