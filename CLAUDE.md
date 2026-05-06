# CLAUDE.md — Cloudflare DNS CLI (cfdns)

## Project Overview

Build a production-grade Go CLI tool named `cfdns` that manages DNS records in Cloudflare.

Current functionality:

* Create A records
* Detect existing records
* Prompt user to update if record exists

Target:
Evolve this into a **fully featured DNS management CLI** with clean architecture, testability, and automation support.

---

## Core Requirements

### CLI Behavior

Command format:

```
cfdns <hostname> <ip>
```

Examples:

```
cfdns test.example.com 192.168.1.1
```

### Environment Variables

* `CLOUDFLARE_API_TOKEN` (required)

Fail fast if missing.

---

## Functional Requirements

### 1. Record Handling

#### Existing behavior (must preserve):

* If record does NOT exist → create it
* If record exists:

  * If same IP → exit cleanly
  * If different IP → prompt user:

    ```
    Update record with new IP? (y/N)
    ```
  * If confirmed → update record

---

### 2. Add CLI Flags

Implement proper flag parsing (prefer `cobra`):

Required flags:

```
--type        (default: A)
--ttl         (default: auto)
--proxied     (default: false)
--force       (skip confirmation prompts)
--json        (machine-readable output)
```

---

### 3. Multi-Command Structure

Refactor into subcommands:

```
cfdns create <hostname> <ip>
cfdns update <hostname> <ip>
cfdns delete <hostname>
cfdns list <domain>
```

---

### 4. Record Types

Support:

* A
* AAAA
* CNAME
* TXT

Auto-detect:

* IPv4 → A
* IPv6 → AAAA

---

### 5. Zone Detection

Replace naive parsing with:

```
golang.org/x/net/publicsuffix
```

Must correctly handle:

* example.com
* example.co.uk
* sub.sub.example.com

---

### 6. Output Modes

#### Human-readable (default)

```
Created: test.example.com -> 192.168.1.1
```

#### JSON (`--json`)

```
{
  "action": "create",
  "hostname": "test.example.com",
  "ip": "192.168.1.1",
  "status": "success"
}
```

---

### 7. Exit Codes

Standardize:

| Code | Meaning       |
| ---- | ------------- |
| 0    | Success       |
| 1    | General error |
| 2    | Invalid input |
| 3    | API error     |
| 4    | No change     |

---

## Architecture Requirements

### Project Structure

```
/cmd
  root.go
  create.go
  update.go
  delete.go
  list.go

/internal
  /cloudflare
    client.go
    dns.go
    zones.go

  /cli
    prompt.go
    output.go

  /util
    validation.go
    domain.go

main.go
```

---

### Design Principles

* No global state
* Dependency injection for API client
* Clear separation:

  * CLI layer
  * API layer
  * business logic
* All Cloudflare logic isolated in `/internal/cloudflare`

---

## Cloudflare Integration

Use:

```
github.com/cloudflare/cloudflare-go
```

Must implement:

* Get Zone ID
* List DNS records
* Create record
* Update record
* Delete record

---

## Validation Rules

### Hostname

* Must be valid FQDN

### IP

* Must validate:

  * IPv4
  * IPv6

Fail early if invalid.

---

## Prompting Rules

* Default: interactive
* If `--force` is set → NEVER prompt
* Prompts must be safe for scripting environments

---

## Logging

* Use structured logging (optional: `log/slog`)
* Do NOT print logs in JSON mode

---

## Testing Requirements

Create:

* Unit tests for:

  * zone extraction
  * IP validation
  * record comparison

* Mock Cloudflare API layer

---

## Future Enhancements (Design For)

* Bulk operations (CSV input)
* Import/export DNS zones
* Watch mode (sync DNS periodically)
* n8n / automation integration
* Config file support

---

## Constraints

* Go 1.22+
* No dotenv files
* Use environment variables only
* Must compile to single static binary
* No unnecessary dependencies

---

## Deliverables

Claude must produce:

1. Fully refactored CLI using cobra
2. Clean package structure
3. Working commands:

   * create
   * update
   * delete
   * list
4. Input validation
5. JSON output mode
6. Proper error handling
7. Buildable binary

---

## Build Instructions

```
go mod tidy
go build -o cfdns
```

---

## Success Criteria

* CLI behaves predictably in both interactive and automation contexts
* Code is modular and extensible
* No breaking changes to existing behavior
* Handles real-world DNS edge cases correctly

---

## Notes for Claude Code

* Do NOT rewrite working logic unnecessarily
* Extend incrementally
* Prioritize correctness over cleverness
* Keep functions small and testable
* Avoid overengineering

---

End of specification.

