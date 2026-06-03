# Go Development Environment Setup

This guide will help you set up a complete Go development environment on a new machine.

## Table of Contents
1. [Install Go](#install-go)
2. [Install VS Code Extensions](#install-vs-code-extensions)
3. [Install Go Development Tools](#install-go-development-tools)
4. [Set Up Your Project](#set-up-your-project)
5. [Building and Running](#building-and-running)

---

## Install Go

### On macOS
```bash
brew install go
```

### On Ubuntu/Debian
```bash
sudo apt update
sudo apt install golang-go
```

### On Windows
Download the installer from https://golang.org/dl/ and run it.

### Verify Installation
```bash
go version
```

---

## Install VS Code Extensions

Open VS Code and install the official Go extension:
- Extension: **Go** (by Go Team at Google)
- Publisher ID: `golang.go`

Or run:
```bash
code --install-extension golang.go
```

---

## Install Go Development Tools

After Go is installed, run these commands to install essential development tools:

```bash
# Debugger - Step through your code
go install github.com/go-delve/delve/cmd/dlv@latest

# Linter - Catch code style issues
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Language Server - Intellisense, hover info, refactoring
go install golang.org/x/tools/gopls@latest

# Static Analyzer - Additional code quality checks
go install honnef.co/go/tools/cmd/staticcheck@latest
```

### Verify All Tools
```bash
dlv version
golangci-lint version
gopls version
staticcheck -version
```

---

## Set Up Your Project

### Initialize a Go Module
```bash
go mod init <module-name>
# Example:
go mod init ch3
```

This creates a `go.mod` file for dependency management.

### Install Dependencies
```bash
go mod tidy
```

---

## Building and Running

### Run Your Program
```bash
go run program.go hello
```

Output:
```
You provided the argument: hello
```

### Build an Executable
```bash
go build -o program
./program hello
```

### Run with Make
```bash
make build
make run
```

### Debug Your Code
```bash
dlv debug program.go
# Inside debugger:
(dlv) break main.main
(dlv) continue
(dlv) next
(dlv) print arg
(dlv) exit
```

### Lint Your Code
```bash
golangci-lint run
```

### Check for Static Issues
```bash
staticcheck ./...
```

---

## Quick Reference

| Command | Purpose |
|---------|---------|
| `go run program.go` | Run program immediately |
| `go build -o program` | Compile to executable |
| `go mod init <name>` | Initialize module |
| `go mod tidy` | Clean up dependencies |
| `dlv debug` | Start debugger |
| `golangci-lint run` | Check code style |
| `staticcheck ./...` | Find potential bugs |

---

## Troubleshooting

### Import getting removed on save?
Add to VS Code settings (`.vscode/settings.json`):
```json
{
  "[go]": {
    "editor.formatOnSave": false,
    "editor.codeActionsOnSave": {
      "source.organizeImports": false
    }
  }
}
```

### Tools not found after installation?
Make sure `$GOPATH/bin` is in your `PATH`:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Add this to your shell profile (`.bashrc`, `.zshrc`, etc.) to make it permanent.

---

## Project Structure
```
ch3/
├── go.mod          # Go module file
├── program.go      # Main program
├── Makefile        # Build automation
└── README.md       # This file
```

---

## Next Steps

1. Modify `program.go` to add new features
2. Run `go run program.go` to test
3. Use `dlv debug` if you need to debug
4. Run `golangci-lint run` before committing
5. Build with `go build` for production
