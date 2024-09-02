# Oats -- One-at-a-Time To-do's

## About

Oats is a simple CLI-based to-do app which only returns one task at a time.
It is best used to organize lists of tasks with no timeliness. If you tend to end
up with long lists of someday to-do's, drop them into Oats to keep them in order
while clearing them from your mind.

## Building

Needs: git, go

1. Clone the repository: `git clone https://github.com/yamlinson/oats.git ~/git/oats`
2. Build a binary: `go build -o ~/git/bin/oats ~/git/oats`
3. Make sure your Go bin dir is in your path: ``export PATH=$PATH`go env GOPATH`/bin``

## Usage

Oats is made up of three main subcommands:

### `oats add ["list name"] ["item name"] [flags]`

`oats add` has no primary flags. `-h` returns help text.

List-item pairs must be unique in the database, but the same item name can
be used across any number of distinct lists.

### `oats get ["list name"] [flags]`

Run without a flag, `oats get "list name"` will return the oldest item in a list.

Flags:

- -A, --all

Get all items from all lists.

- -a, --all-in-list

Get all items from a given list, or the names of all lists if no list is supplied.

- -R, --any-random

Get a random item from any list.

- -c, --current

Get the most recently returned item.

- -l, --last

Get the most recently created item in a list instead of the oldest.

- -r, --random

Get a random item from the specified list.

### `oats rm ["list name"] ["item name"] [flags]`

Run without a flag, `oats rm "list name" "item name"` removes the given item from the given list.

Flags:

- -c, --current

Remove the most recently returned item.
