# a5proxy

a5proxy is a proxy server that provides OpenAI/Gemini/Claude/Codex/Grok compatible API
interfaces for CLI coding tools, with OAuth logins and round-robin load balancing across
multiple accounts.

a5proxy is a fork of [router-for-me/CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI).
The fork lives at [andershfranzen/CLIProxyAPI](https://github.com/andershfranzen/CLIProxyAPI).
User-visible branding is "a5proxy"; internal identifiers (Go module path, package names,
API paths, config keys) intentionally keep the upstream names so upstream changes merge easily.

## Features

- OpenAI (Chat Completions and Responses), Gemini (including Interactions), and Claude compatible endpoints
- OAuth logins for Codex, Claude Code, Grok Build, Antigravity, Kimi, Devin and more
- Streaming, non-streaming and WebSocket responses where supported
- Function calling/tools and multimodal input
- Multiple accounts with round-robin load balancing
- OpenAI-compatible upstream providers via config (e.g. OpenRouter)
- Management API (`/v8/management`) and a bundled web management panel (`/management.html`)
- Terminal UI (`--tui`, `--standalone`)
- Reusable Go SDK for embedding the proxy

## Build and run

```bash
go build -o a5proxy ./cmd/server
cp config.example.yaml config.yaml   # then edit api-keys, management secret, etc.
./a5proxy --config config.yaml
```

Useful flags: `--config <path>`, `--tui`, `--standalone`, `--local-model`, `--no-browser`,
`--oauth-callback-port <port>`, and the provider login flags (`--codex-login`, `--claude-login`, ...).
Run `./a5proxy -h` for the full list.

Credentials are stored in `oauth.auth-dir` (default `~/.a5proxy`). An explicitly configured
`auth-dir` is always used as-is.

## Docker

The image bundles the management web panel, built from the
[management panel](https://github.com/andershfranzen/Cli-Proxy-API-Management-Center) source passed
as the named build context `panel`. The server never downloads the panel at runtime.

```bash
docker build \
  --build-context panel=/path/to/Cli-Proxy-API-Management-Center \
  --build-arg VERSION="$(git describe --tags --always --dirty)" \
  --build-arg COMMIT="$(git rev-parse --short HEAD)" \
  --build-arg BUILD_DATE="$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  -t a5proxy:latest .
```

In the container the binary and config live in `/a5proxy` (`/a5proxy/config.yaml`), the panel is
served from `/a5proxy/static/management.html`, and the default auth directory is `/root/.a5proxy`.
See `docker-compose.yml` for an example deployment (`A5PROXY_*` variables; the legacy `CLI_PROXY_*`
names still work).

## Documentation

- Configuration reference: [config.example.yaml](config.example.yaml)
- Management API: [docs/management-api-v8.md](docs/management-api-v8.md)
- LAN discovery: [docs/lan-discovery.md](docs/lan-discovery.md)
- SDK: [usage](docs/sdk-usage.md), [advanced](docs/sdk-advanced.md), [access](docs/sdk-access.md), [watcher](docs/sdk-watcher.md)
- Examples: [Claude Code with a Codex backend](examples/claude-code-codex/README.md), [plugins](examples/plugin/README.md), `examples/custom-provider`

## License

MIT, see [LICENSE](LICENSE). Based on CLIProxyAPI by Router-For.ME.
