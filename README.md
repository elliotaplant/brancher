# Brancher

A tool for making branches that track issues on GitHub. Requires [hub](https://github.com/github/hub)

## What it does

Makes feature/bug branches and personal branches off of those features branches with a helpful dropdown menu

![Demo](./demo.gif)

## Installation

- Go to the [releases](https://github.com/elliotaplant/brancher/releases) page and download the binary for your OS
- Make the binary an executable (`chmod +x /path/to/brancher-your-os-name`)
- Move the brancher somewhere into your `$PATH` (maybe `/usr/local/bin`)
- Add `alias brancher='brancher-<your-os> <your-initiaals>'` to your `.bashrc` or `.zshrc` file
