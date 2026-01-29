# gorandom

CLI tool for learning Go through curated articles from official sources.

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

## Stack

- [Bun](https://bun.sh) - Runtime
- [citty](https://github.com/unjs/citty) - CLI framework
- [@inquirer/prompts](https://github.com/SBoudrias/Inquirer.js) - Interactive prompts
- [Biome](https://biomejs.dev) - Linter/formatter
