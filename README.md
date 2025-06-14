<<<<<<< HEAD
# bd-pokedexcli
=======
# BDxPokedex CLI

A command-line interface tool for accessing and exploring Pokemon data. Built as a portfolio project in Go.

## Requirements

- Go 1.24 or later

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/print-agent/bd-pokedexcli.git
cd bd-pokedexcli

# Build the application
go build -o bdxpokedex ./cmd/bdxpokedex

# Run the application
./bdxpokedex
```

## Project Structure

- `cmd/bdxpokedex/`: Contains the main application entry point
- `pkg/`: Contains public packages that can be imported by other projects
- `internal/`: Contains private application code not meant to be imported by other projects

## Usage

Currently, the application displays a version message. Additional features will be added in future releases.

```bash
# Run the CLI tool
./bdxpokedex
```

## License

[MIT License](LICENSE)

>>>>>>> ff76e7d (Initial project skeleton for bd-pokedexcli)
