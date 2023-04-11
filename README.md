# gh-organization-webhooks

A GitHub `gh` [CLI](https://cli.github.com/) extension to create a report containing Webhooks defined at an Organization level. The `csv` report includes:

* `Type`: Indicates that the Webhook was created at the `Organization` level.
* `ID`: Associated `id` for the webhook.
* `Name`: Must be passed as `web` if created from the API. Name can only be set to `web` or `email`.
* `Active`: If notifications are sent when the webhook is triggered. Set to `true` to send notifications.
* `Events`: What events the hook is triggered for. Set to `["*"]` to receive all possible events. The default is `["push"]`.
* `Config_ContentType`: The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
* `Config_InsecureSSL`: Whether the SSL certificate of the host for url is verified when delivering payloads. Supported values include `0` (verification is performed) and `1` (verification is not performed). The default is `0`.
* `Config_Secret`: If a `secret` was present when the Webhook was created, the report will return `********`.
* `Config_URL`: The URL to which the payloads are delivered.
* `Updated_At`: Date that the webhook was last updated.
* `Created_At`: Date that the webhook was created.

> *NOTE:*
> This extension does NOT retrieve the value of the webhook secret, and only identifies that one was created.

## Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation) instructions.

2. Install the extension:
  ```sh
  gh extension install katiem0/gh-organization-webhooks
  ```

For more information: [`gh extension install`](https://cli.github.com/manual/gh_extension_install).

## Usage

This extension supports listing and creating webhooks between `GitHub.com` and GitHub Enterprise Server, through the use of `--hostname` and `--source-hostname`.

```sh
$ gh organization-webhooks -h 
List and create organization level webhooks.

Usage:
  organization-webhooks [command]

Available Commands:
  create      Create organization level webhooks
  list        List organization level webhooks

Flags:
  -h, --help   help for organization-webhooks

Use "organization-webhooks [command] --help" for more information about a command.
```

### List Webhooks

```sh
$ gh organization-webhooks list -h
List organization level webhooks

Usage:
  organization-webhooks list <source organization> [flags]

Flags:
  -d, --debug                To debug logging
  -h, --help                 help for list
      --hostname string      GitHub Enterprise Server hostname (default "github.com")
  -o, --output-file string   Name of file to write CSV list to (default "WebhookReport-20230411160920.csv")
  -t, --token string         GitHub personal access token for reading source organization (default "gh auth token")
```


### Create Webhooks

```sh
$ gh organization-webhooks create -h
Create organization level webhooks

Usage:
  organization-webhooks create <target organization> [flags]

Flags:
  -d, --debug                        To debug logging
  -f, --from-file string             Path and Name of CSV file to create webhooks from
  -h, --help                         help for create
      --hostname string              GitHub Enterprise Server hostname (default "github.com")
      --source-hostname string       GitHub Enterprise Server hostname where webhooks are copied from (default "github.com")
  -o, --source-organization string   Name of the Source Organization to copy webhooks from (Requires --source-token)
  -s, --source-token string          GitHub personal access token for Source Organization (Required for --source-organization)
  -t, --token string                 GitHub personal access token for organization to write to (default "gh auth token")
```