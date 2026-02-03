# gorandom

A CLI tool to discover random Go articles, tutorials, and documentation from trusted sources.

![Golang](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)

## Features

- **Random Article**: Get a random article to read.
- **Source Filtering**: Filter by docs, tour, Go by Example, stdlib pkg, or blog.
- **History**: Keep track of what you've read.
- **Progress Tracking**: See your reading stats and progress bars.
- **Interactive Mode**: Mark articles as read/unread using an interactive list.
- **Smart Filtering**: By default, shows only unread articles (unless you've read them all).
- **Next**: Pick the next unread article with a clear message when you're done.
- **Search**: Find articles by title or source.
- **Non-Interactive Output**: `--plain`, `--json`, or `--no-ui` for scripting.
- **Clipboard**: Copy the current URL from the TUI.

## Installation

### From Source

Requirements: Go 1.21 or later.

```bash
git clone https://github.com/AfshinJalili/gorandom.git
cd gorandom
go install ./cmd/gorandom
```

This will install the `gorandom` binary to your `$GOPATH/bin`. Ensure this directory is in your `PATH`.

### Go Install (Recommended)

```bash
go install github.com/AfshinJalili/gorandom/cmd/gorandom@latest
```

After installing, fetch the latest sources:

```bash
gorandom sources update
```

### Download Binary (GitHub Releases)

1. Download the binary for your OS/arch from the Releases page.
2. Make it executable (macOS/Linux): `chmod +x gorandom`
3. Move it into your `PATH` (example): `mv gorandom /usr/local/bin/`

## Usage

### Basic Commands

```bash
# Get a random unread article
gorandom

# Get a random article from a specific source
gorandom --source tour
gorandom -s blog

# Include already-read articles in the random pick
gorandom --any

# View your recent history
gorandom history
gorandom history --limit 20

# Pick the next unread article
gorandom next

# Search by keyword (title/source)
gorandom search generics
```

### Learning Tools

```bash
# Open article in browser
gorandom open
gorandom open 1

# Manage Bookmarks
gorandom bookmark 1    # Toggle bookmark
gorandom bookmarks     # List bookmarks

# Check progress & streaks
gorandom stats
```

### Non-Interactive Output

```bash
# Plain text
gorandom random --plain
gorandom history --plain --limit 5

# JSON
gorandom random --json
gorandom history --json --limit 20

# Disable UI (alias of --plain)
gorandom random --no-ui
```

### TUI Shortcuts

- `n` / space: next article
- `o` / enter: open in browser
- `m`: mark as read
- `b`: bookmark
- `y`: copy URL
- `h`: toggle help
- `q`: quit
### Tracking Progress

```bash
# Mark an article as read (interactive)
gorandom mark

# Mark specific URL as read
gorandom mark https://go.dev/tour/welcome/1

# Mark most recently viewed article as read
gorandom mark 1

# Mark as unread (interactive or by argument)
gorandom unmark

# View your learning stats
gorandom stats
```

## Reviewing Sources

```bash
# List available content sources and article counts
gorandom sources
```

## Development

The project is written in Go and uses:
- [Cobra](https://github.com/spf13/cobra) for CLI commands
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for TUI components

### Running Locally

```bash
go run ./cmd/gorandom
```

### Running Tests

```bash
go test ./...
```

### Linting and Formatting

```bash
make fmt
make lint
make vet
make race
```

### Git Hooks (Pre-Commit)

```bash
./scripts/install-hooks.sh
```

The pre-commit hook auto-runs `gofmt` and `golangci-lint`.

### Release (GoReleaser)

```bash
make release-dry
```

### Publishing a Release (Step by Step)

```bash
# 1) Ensure you're on main and clean
git status

# 2) Run tests (recommended)
go test ./...
make fmt
make lint

# 3) Tag the release (use semantic versioning)
git tag v0.1.1

# 4) Push code and tag
git push origin main
git push origin v0.1.1
```

After pushing the tag, GitHub Actions runs GoReleaser to build binaries
for Linux/macOS/Windows and publishes a GitHub Release with the artifacts.

Data storage defaults to `~/.random-go/history.json`.
For testing, you can override the config directory:

```bash
GORANDOM_CONFIG_DIR=/tmp/test-go go run ./cmd/gorandom
```
