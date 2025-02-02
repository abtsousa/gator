
# Gator - A command-line RSS aggregator

Gator is a simple command-line RSS feed aggregator written in Go. It allows you to follow RSS feeds and browse posts from your favorite websites.

## Prerequisites

- Go 1.21 or later
- PostgreSQL 14 or later

## Installation

1. Install the CLI using Go:
```bash
go install github.com/abtsousa/gator@latest
```

2. Create a configuration file at `~/.gatorconfig.json` with your Postgres connection string:
```json
{
    "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

3. Set up the database schema using [goose](https://github.com/pressly/goose) or by running the SQL files in `sql/schema/` manually.

```bash
cd sql/schema
goose postgres <db_url> down
goose postgres <db_url> up
```

## Usage

Here are some of the available commands:

- `gator register <username>` - Create a new user account
- `gator login <username>` - Log in as an existing user
- `gator users` - List all users
- `gator addfeed <name> <url>` - Add a new RSS feed
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow an existing feed
- `gator following` - List feeds you're following
- `gator unfollow <url>` - Unfollow a feed
- `gator browse [limit]` - Browse posts from your feeds (default: 2 posts)

Example usage:
```bash
gator register johndoe
gator addfeed "Example Blog" "https://example.com/feed.xml"
gator browse 5
```

You can also run the aggregator in the background to fetch new posts periodically:
```bash
gator agg 1m  # Fetch every minute
```
