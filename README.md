# notifier

[![go][go-version-src]][go-version-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

A simple notifier CLI for Discord and Slack with Webhook. Powered by [Go][go-version-href].

- `webhook`: send a message to Discord and Slack from CLI arguments.
- `send`: send a message to Discord and Slack from `.env` credentials.

## Install

From Go (you cannot use `send` command with this method, you need to use `webhook` command)

```bash
go get github.com/ewilan-riviere/notifier
```

From source

```bash
git clone https://github.com/ewilan-riviere/notifier.git
cd notifier
go build -o notifier
```

Create symlink to use `notifier` everywhere

```bash
sudo rm -f /usr/local/bin/notifier
sudo ln -s ~/notifier/notifier /usr/local/bin
sudo chmod +x /usr/local/bin/notifier
```

## Usage

### From CLI arguments

Send a message to Discord or Slack.

```bash
notifier webhook https://discord.com/api/webhooks/XXXXXXXXX/XXXXXXXXX "Hello World"
notifier webhook https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXX "Hello World"
```

### From `.env` credentials

> [IMPORTANT]
>
> Available only if you install from source.

Create `.env` file.

```bash
cp .env.example .env
```

Add your Discord and Slack webhooks, you can leave it blank.

```yaml
DISCORD_WEBHOOK=XXXXXXXXX/XXXXXXXXX
SLACK_WEBHOOK=XXXXXXXXX/XXXXXXXXX/XXXXXXXXX
```

Send a message to Discord or Slack.

```bash
notifier send discord "Hello World"
notifier send slack "Hello World"
```

## Example with GitLab CI

You can install locally `notifier` and use it in your GitLab CI. The better solution is to use `.env` file, with installation from source. But in this example, we use `webhook` command.

You can create a script, for example `notify`, on your server, on `/usr/local/bin` for example to use it everywhere.

```bash
#!/bin/bash

notifier webhook https://discord.com/api/webhooks/XXXXXXXXX/XXXXXXXXX "$1"
notifier webhook https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXX "$1"
```

Allow execution of this script.

```bash
chmod +x /usr/local/bin/notify
```

Use it in your GitLab CI, without `.env` file and hidden credentials.

```yaml
stages:
  - deploy

deploy-job:
  stage: deploy
  image: alpine:latest
  script:
    - notify '$CI_PROJECT_TITLE deployed'
  only:
    - main
```

## Test

To locally test

```bash
go build -o notifier && ./notifier
```

Execute tests

```bash
go test ./pkg/... -coverprofile=coverage.out
go test -v ./...
go test -v ./pkg/webhook
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[tests-src]: https://img.shields.io/github/actions/workflow/status/kiwilan/php-rss/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://packagist.org/packages/kiwilan/php-rss
[license-src]: https://img.shields.io/github/license/ewilan-riviere/notifier.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/notifier/blob/main/LICENSE
