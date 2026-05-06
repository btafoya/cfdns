# cfdns

[![Go Version](https://img.shields.io/github/go-mod/go-version/btafoya/cfdns?style=flat)](https://golang.org)
[![License](https://img.shields.io/github/license/btafoya/cfdns?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/btafoya/cfdns?style=flat)](https://goreportcard.com/report/github.com/btafoya/cfdns)
[![GitHub Release](https://img.shields.io/github/v/release/btafoya/cfdns?style=flat)](https://github.com/btafoya/cfdns/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/btafoya/cfdns/build.yml?style=flat)](https://github.com/btafoya/cfdns/actions)
[![Stars](https://img.shields.io/github/stars/btafoya/cfdns?style=flat)](https://github.com/btafoya/cfdns/stargazers)
[![Forks](https://img.shields.io/github/forks/btafoya/cfdns?style=flat)](https://github.com/btafoya/cfdns/network/members)

A fast, simple CLI for managing Cloudflare DNS records. Create, update, delete, and list records from your terminal — with JSON output for automation pipelines.

---

## Features

- Create, update, delete, and list DNS records
- Auto-detects A vs AAAA from IPv4/IPv6 input
- Smart zone detection (handles `example.co.uk`, `sub.sub.example.com`, etc.)
- Supports A, AAAA, CNAME, and TXT record types
- JSON output mode for scripting and automation
- `--force` flag for non-interactive use in CI/CD
- Confirms before destructive operations (unless `--force`)
- Structured exit codes for reliable scripting

---

## Install

Download the latest release from [GitHub Releases](https://github.com/btafoya/cfdns/releases/latest).

### Linux — apt (Debian/Ubuntu)

```bash
wget https://github.com/btafoya/cfdns/releases/latest/download/cfdns_<version>_amd64.deb
sudo dpkg -i cfdns_<version>_amd64.deb
```

### Linux — rpm (Fedora/RHEL/CentOS)

```bash
sudo rpm -i https://github.com/btafoya/cfdns/releases/latest/download/cfdns-<version>-1.x86_64.rpm
```

### Linux — AppImage

```bash
wget https://github.com/btafoya/cfdns/releases/latest/download/cfdns-<version>-x86_64.AppImage
chmod +x cfdns-<version>-x86_64.AppImage
./cfdns-<version>-x86_64.AppImage
```

### macOS

```bash
# Apple Silicon (M1/M2/M3)
curl -L https://github.com/btafoya/cfdns/releases/latest/download/cfdns-<version>-darwin-arm64.tar.gz | tar -xz
sudo mv cfdns-darwin-arm64 /usr/local/bin/cfdns

# Intel
curl -L https://github.com/btafoya/cfdns/releases/latest/download/cfdns-<version>-darwin-amd64.tar.gz | tar -xz
sudo mv cfdns-darwin-amd64 /usr/local/bin/cfdns
```

### Windows

Download `cfdns-<version>-windows-amd64.zip` from [releases](https://github.com/btafoya/cfdns/releases/latest), extract, and add the directory to your `PATH`.

### go install

```bash
go install github.com/btafoya/cfdns@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

### Build from source

```bash
git clone https://github.com/btafoya/cfdns.git
cd cfdns
go build -o cfdns
```

---

## Setup

Export your Cloudflare API token:

```bash
export CLOUDFLARE_API_TOKEN=your_token_here
```

The token needs **Zone:DNS:Edit** permissions for the target zone.

---

## Usage

```
cfdns <command> [flags] [args]
```

### Commands

| Command | Description |
|---------|-------------|
| `create <hostname> <content>` | Create a DNS record |
| `update <hostname> <content>` | Update an existing record |
| `delete <hostname>` | Delete a record |
| `list <domain>` | List all records for a domain |

### Global Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--type` | `A` | Record type: A, AAAA, CNAME, TXT |
| `--ttl` | `1` | TTL in seconds (1 = auto) |
| `--proxied` | `false` | Enable Cloudflare proxy (orange cloud) |
| `--force` | `false` | Skip confirmation prompts |
| `--json` | `false` | Output JSON instead of human-readable text |

---

## Examples

### Create a record

```bash
cfdns create test.example.com 192.168.1.1
# Created: test.example.com -> 192.168.1.1
```

### Auto-detect IPv6

```bash
cfdns create test.example.com 2001:db8::1
# Creates AAAA record automatically
```

### Create a CNAME

```bash
cfdns create www.example.com example.com --type CNAME
```

### List all records

```bash
cfdns list example.com
# A      test.example.com    192.168.1.1
# CNAME  www.example.com     example.com
```

### Delete a record

```bash
cfdns delete test.example.com
# Delete A test.example.com -> 192.168.1.1? (y/N):
```

### Force delete (no prompt)

```bash
cfdns delete test.example.com --force
```

---

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | General error |
| `2` | Invalid input |
| `3` | API error |
| `4` | No change needed |

---

## Contributing

Contributions are welcome! Please open an issue before submitting a PR for significant changes.

1. Fork the repo
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit your changes
4. Push and open a Pull Request

---

## Author

Built by [Brian Tafoya](https://briantafoya.com)

---

## License

[MIT](LICENSE) — © Brian Tafoya
