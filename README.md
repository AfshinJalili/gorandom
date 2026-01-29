# random-go

A CLI tool that serves random Go learning articles from curated sources.

## Sources

- **Go Docs** - Official documentation and tutorials
- **Tour of Go** - Interactive Go tutorial
- **Go by Example** - Annotated example programs
- **Standard Library** - Package documentation (pkg.go.dev)
- **Go Blog** - Official Go blog posts

## Installation

```bash
# Clone and install
bun install

# Run directly
bun run index.ts

# Or link globally
bun link
random-go
```

## Usage

```bash
# Get a random unread article
random-go

# Get any random article (including ones you've read)
gorandom--any

# Filter by source
gorandom--source tour
gorandom--source blog
gorandom--source gobyexample

# View your history
gorandomhistory
gorandomhistory --limit 20

# Mark article as read (use index 1 for most recent)
gorandommark 1
gorandommark https://gobyexample.com/goroutines

# Unmark article
gorandomunmark 1

# See your progress
gorandomstats

# List available sources
gorandomsources

# Show help
gorandomhelp
```

## Features

- **Random Selection**: Prioritizes unread articles by default
- **History Tracking**: Automatically records viewed articles
- **Mark as Read**: Track your progress through Go learning materials
- **Source Filtering**: Focus on specific content sources
- **Progress Stats**: See completion percentages per source

## Data Storage

History is stored in `~/.random-go/history.json`

## Adding More Links

Edit `links.ts` to add more articles to the pool:

```typescript
{ url: 'https://example.com/article', source: 'blog', title: 'Article Title' }
```

Valid sources: `docs`, `tour`, `gobyexample`, `pkg`, `blog`
