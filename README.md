Functions:
- Emoji stat tracking
- Rename people at the top of the member list with the non-alpha character names
- Emoji stealing with PNG

Soon:
- Emoji deletion by ID
- Add buttons and fancy colors and emoji pngs to the site

If you want to invite the bot to your discord server:
https://discord.com/oauth2/authorize?client_id=876719578858811392&scope=bot&permissions=413592308800

Emoji stats site: http://www.robinai.xyz/

Installation methods differ between operating systems, only do if you plan to self host:

Install GoLang
Install PHP, latest preferrably
Install nginx (or xampp). Configure to localhost/site if you have a site
(The frontend files are under the html folder, you can either redirect the nginx config or just move the html to the default folder)
Install MySQL/MariaDB, and set it up, make a database 
Put the Database credentials over in robinAI/bot/db.go, and in html/includes/dbh.inc.php
Put your discord bot token over in robinAI/config/config.json
open a terminal, cd <path-to-robinai>, go build, go run .
