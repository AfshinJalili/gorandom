# gorandom

A CLI and TUI to discover random Go articles, tutorials, and documentation from trusted sources.

![Golang](https://img.shields.io/badge/Go-1.25.5+-00ADD8?style=flat&logo=go)

## Features

- Random article picker with unread-first behavior
- Source filtering across docs, tour, Go by Example, stdlib pkg, and blog
- History, bookmarks, and streak tracking
- Interactive TUI with keyboard shortcuts
- Non-interactive output for scripting (`--plain`, `--json`, `--no-ui`)
- Remote sources file with local caching and manual refresh

## Installation

### Go Install (Recommended)

```bash
go install github.com/AfshinJalili/gorandom/cmd/gorandom@latest
```

### From Source

```bash
git clone https://github.com/AfshinJalili/gorandom.git
cd gorandom
go install ./cmd/gorandom
```

### Download Binary (GitHub Releases)

1. Download the binary for your OS and architecture from GitHub Releases.
2. Make it executable (macOS/Linux): `chmod +x gorandom`
3. Move it into your `PATH`, for example: `mv gorandom /usr/local/bin/`

## Usage

### Core Commands

```bash
# Random unread article
gorandom

# Pick next unread article
gorandom next

# Search by keyword (title or source)
gorandom search generics

# Filter by source
gorandom --source tour
gorandom next --source blog

# Include already-read articles
gorandom --any
```

### History and Bookmarks

```bash
# Recent history
gorandom history
gorandom history --limit 20

# Mark read/unread
gorandom mark
gorandom unmark

# Bookmark
gorandom bookmark 1
gorandom bookmarks

# Progress and streak
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

### Sources and Cache

```bash
# Fetch or refresh sources
gorandom sources update

# Cache status
gorandom sources status
```

## TUI Shortcuts

- `n` or space: next article
- `o` or enter: open in browser
- `m`: toggle read
- `b`: toggle bookmark
- `y`: copy URL
- `h`: toggle help
- `q`: quit

## Data and Caching

Sources are loaded from a JSON file hosted in this repo:

```
https://raw.githubusercontent.com/AfshinJalili/gorandom/main/data/sources.json
```

Behavior:

- On first run, the CLI fetches the sources file and caches it locally.
- The cache lives in the config directory: `~/.random-go/sources.json`.
- Use `gorandom sources update` to refresh.
- `gorandom sources status` shows cache age and staleness (24h TTL indicator).

## Configuration

Environment variables:

- `GORANDOM_CONFIG_DIR`: override config directory
- `GORANDOM_SOURCES_URL`: override sources JSON URL
- `GORANDOM_SOURCES_TTL`: override stale threshold used by status (default `24h`)
- `GORANDOM_SOURCES_SPINNER=0`: disable the fetch spinner

## Development

### Run Locally

```bash
go run ./cmd/gorandom
```

### Tests

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

## Troubleshooting

- **Sources cache missing**: run `gorandom sources update`.
- **Clipboard on Linux**: install `xclip` or `xsel` to enable copying.
- **History file location**: `~/.random-go/history.json` (or `GORANDOM_CONFIG_DIR`).
