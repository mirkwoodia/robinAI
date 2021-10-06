A bot made by Mirk#6460 (heres my ID, send me a dm maybe say hi 280713802356490241)

Functions:
- Local Emoji stat tracking
- Option to rename the non-alpha character names (using r.noname)

Soon:
- Emoji deletion by ID
- Emoji stealing with PNG
- Check messages for virus, with options if virus possibility detected
- Add buttons and fancy colors and emoji pngs to the site

If you want to invite the bot to your discord server:
https://discord.com/oauth2/authorize?client_id=876719578858811392&scope=bot&permissions=413592308800

Emoji stats site: http://www.robinai.xyz/

Installation methods differ between operating systems, requirements:

Install GoLang
Install PHP, latest preferrably
Install nginx (or xampp). Configure to localhost/site if you have a site
(The frontend files are under the html folder, you can either redirect the nginx config or just move the html to the default folder)
Install MySQL/MariaDB, and set it up, make a database 
Put the Database credentials over in robinAI/bot/db.go, and in html/includes/dbh.inc.php
Put your discord bot token over in robinAI/config/config.json
open a terminal, cd <path-to-robinai>, go build, go run .