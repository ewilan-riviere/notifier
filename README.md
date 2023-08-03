# Notifier

[![go][go-version-src]][go-version-href]
[![license][license-src]][license-href]

[go-version-src]: https://img.shields.io/static/v1?style=flat-square&label=Go&message=v1.20&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[license-src]: https://img.shields.io/github/license/kiwilan/php-rss.svg?style=flat-square&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/kiwilan/php-rss/blob/main/README.md

## Install

Create `.env` file.

```bash
cp .env.example .env
```

Add your Discord and Slack webhooks (you can add multiple servers or leave it blank)

```yaml
DISCORD_SERVERS=default:XXXXXXXXX:XXXXXXXXX,custom:XXXXXXXXX:XXXXXXXXX
SLACK_SERVERS=default:XXXXXXXXX:XXXXXXXXX:XXXXXXXXX,custom:XXXXXXXXX:XXXXXXXXX:XXXXXXXXX
```

Install dependencies.

```bash
go get github.com/joho/godotenv
```

Build the script.

```bash
go build -o notifier
```

You can use `./notifier` to run the script.

```bash
./notifier hello
```

## Global usage

```bash
sudo rm -f /usr/local/bin/notifier
```

```bash
sudo ln -s /home/kiwilan/notifier/notifier /usr/local/bin
```

```bash
sudo chmod +x /usr/local/bin/notifier
```

```bash
notifier hello
```

## Test

Check with `curl` if the webhook is working.

```bash
curl -i -H "Accept: application/json" \
  -H "Content-Type:application/json" \
  -X POST --data \
  "{\"content\": \"Posted Via Command line\"}" \
  https://discord.com/api/webhooks/SERVER_ID/SERVER_TOKEN
```
