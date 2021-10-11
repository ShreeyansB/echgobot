# echgobot
A simple bot that ECHO's stuff.

## Commands
```
/start - Info
/echo - Echo stuff normally
/echob - Echo stuff in bold
/echoi - Echo stuff in italic
/echou - Echo stuff underlined
/echom - Echo stuff in monospace
/echo(b/i/u) - Echo stuff in combinations of b/i/u, order alphabetically
```

## Deploying

### Bot Setup 
Create a bot [@BotFather](https://t.me/botfather). Get the API Key.
Currently uses Env Variable to whitelist only one group to the bot.
Change the code on Line 22 and Line 86 to allow all groups/chats.

Get the group/chat it and set it as:
```
BOT_TOKEN=14532169:JksmKMS975SAYoUrT0k3NdqiqNS84nSUD89e
GROUP_ID=-1039694202194
```
in a new file called .env in root dir.

Finally, set the webhook for your server by CURL-ing the Telegram API.

This can be done by running the following cURL command :

```
curl https://api.telegram.org/bot<token>/setWebhook \
    -F "url=<forwarding url>"
```

Replace ```token``` with your Bot Token, and ```forwarding url``` with your webhook url.

Example :
```
curl https://api.telegram.org/bot14532169:JksmKMS975SAYoUrT0k3NdqiqNS84nSUD89e/setWebhook \
    -F "url=https://ce8b-1115-111-1109-1179-1130-1124-1187-117e.ngrok.io"
```

Once all of this is done, run the bot by running the command :

```
go run main.go
```

<img src="ss.jpg" width="500">

> Also use **go mod vendor** to download libraries in root folder for deployment. I used Heroku for deployment. Dont forget to set the environment variables for your preffered method of deployment. Eg: Set Config Vars in Heroku Dashboard.


