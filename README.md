# github-activity

A CLI tool that fetches and displays a GitHub user's recent public activity using the GitHub Events API.

## Prerequisites

- [Go](https://golang.org/dl/) 1.21 or later

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/sakaicodes/github-activity.git
cd github-activity
go build
```

## Usage

```bash
./github-activity [--filter <event-type>] <github-username>
```

### Arguments

| Argument | Description |
|---|---|
| `<github-username>` | **Required.** The GitHub username to fetch activity for. |
| `--filter <event-type>` | Optional. Filter results by a specific event type. |

### Common Event Types

| Event Type | Description |
|---|---|
| `PushEvent` | Commits pushed to a repository |
| `PullRequestEvent` | Pull request opened, closed, or merged |
| `IssuesEvent` | Issue opened or closed |
| `WatchEvent` | Repository starred |
| `ForkEvent` | Repository forked |
| `CreateEvent` | Branch or tag created |
| `DeleteEvent` | Branch or tag deleted |

## Examples

Fetch all recent activity for a user:

```bash
./github-activity sakaicodes
```

Fetch only push events for a user:

```bash
./github-activity --filter PushEvent sakaicodes
```

Fetch only pull request events:

```bash
./github-activity --filter PullRequestEvent sakaicodes
```

## Run Without Building

```bash
go run . <github-username>
go run . --filter PushEvent <github-username>
```
CLI for interacting with GitHub activity written in Golang
