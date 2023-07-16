# Notifier

[![go][go-version-src]][go-version-href]
[![license][license-src]][license-href]

[go-version-src]: https://img.shields.io/static/v1?style=flat-square&label=Go&message=v1.20&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[license-src]: https://img.shields.io/github/license/kiwilan/php-rss.svg?style=flat-square&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/kiwilan/php-rss/blob/main/README.md

## Local

```bash
go get github.com/joho/godotenv
```

## Build

```bash
go build -o notifier
./notifier hello
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
