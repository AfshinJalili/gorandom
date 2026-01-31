# gorandom

CLI tool for learning Go through curated articles from official sources.

## Prerequisites

Requires [Bun](https://bun.sh) runtime.

```bash
# Install Bun (macOS, Linux, WSL)
curl -fsSL https://bun.sh/install | bash
```

## Install

```bash
bun install
bun link
```

## Usage

```bash
# Random unread article
gorandom

# Include already-read articles
gorandom --any

# Filter by source
gorandom --source=tour
gorandom random -s blog

# View history
gorandom history
gorandom history --limit=5

# Mark as read (index or URL)
gorandom mark 1
gorandom mark https://go.dev/doc/effective_go

# Interactive mark (no arg = select from list)
gorandom mark

# Unmark as unread
gorandom unmark 1

# Progress stats
gorandom stats

# List sources
gorandom sources

# Help
gorandom --help
gorandom mark --help
```

## Sources (258 articles)

| Source | Description | Count |
|--------|-------------|-------|
| `docs` | Official Go documentation | 26 |
| `tour` | Tour of Go interactive tutorial | 83 |
| `gobyexample` | Go by Example | 82 |
| `pkg` | Standard library docs | 24 |
| `blog` | Official Go Blog | 43 |

## Development

```bash
# Run
bun run start

# Test
bun test

# Lint & format
bun run check
```

## Storage

History saved to `~/.random-go/history.json`

## How commands work

### history

- Loads `~/.random-go/history.json` (creates dir/file if missing).
- Sorts entries by `viewedAt` descending (newest first).
- Shows the first `limit` entries (default 10). Each line: `[READ]` or `[    ]`, date, title, then URL.
- Entries are added/updated when you run `gorandom` (random article) or use `mark`/`unmark`.

### mark (mark as read)

**Resolving the target URL:**

- **No arg:** Interactive prompt lists the 10 most recent **unread** entries; you pick one.
- **Number (e.g. `1`):** History sorted newest-first; `1` = most recent, `2` = second, etc. Invalid index exits with error.
- **Other:** Treated as the URL.

**Effect:** `markAsRead(url)` loads history; if the URL exists, sets `isRead = true`; if not, adds a new entry with `isRead = true`. Always succeeds.

### unmark (mark as unread)

**Resolving the target URL:**

- **No arg:** Interactive prompt lists the 10 most recent **read** entries; you pick one.
- **Number:** Same as mark — 1 = most recent in history. Invalid index exits with error.
- **Other:** Treated as the URL.

**Effect:** `markAsUnread(url)` only updates existing entries (sets `isRead = false`). If the URL is not in history, prints "URL not found in history" and does nothing.

## Stack

- [Bun](https://bun.sh) - Runtime
- [citty](https://github.com/unjs/citty) - CLI framework
- [@inquirer/prompts](https://github.com/SBoudrias/Inquirer.js) - Interactive prompts
- [Biome](https://biomejs.dev) - Linter/formatter
