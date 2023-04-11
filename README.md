# gh-organization-webhooks

A GitHub `gh` [CLI](https://cli.github.com/) extension to create a report containing Webhooks defined at an Organization level. The `csv` report includes:

* `Type`
* `ID`
* `Name`
* `Active`
* `Events`
* `Config_ContentType`
* `Config_InsecureSSL`
* `Config_Secret`
* `Config_URL`
* `Updated_At`
* `Created_At`

> *NOTE:*
> This extension does NOT retrieve the value of the webhook secret, and only identifies that one was created.

## Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation) instructions.

2. Install the extension:
  ```sh
  gh extension install katiem0/gh-organization-webhooks
  ```

For more information: []`gh extension install`](https://cli.github.com/manual/gh_extension_install).

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