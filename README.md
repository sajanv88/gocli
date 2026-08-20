# gocli

A command-line tool that scaffolds a working Go backend, wires up a database, and spins up a frontend, all in one command.

## Why this exists

Every time I started a new Go service I did the same five things: run `go mod init`, wire up a router, hand-write a docker-compose file for whatever database I picked that week, remember the right flags for `create-next-app`, and write a `.gitignore` I'd copy from the last project. None of it was hard. All of it was the same fifteen minutes, over and over, before I got to write a single line of actual application code.

gocli does that in under a minute for you. Pick a router, a database, a frontend, and it builds the whole thing: connection code, a docker-compose file, an env example, a README specific to what you picked, and, if you want it, a Dockerfile. The frontend is scaffolded by calling Vite's and Next.js's own official CLIs directly, not by reimplementing them, so you get exactly what those tools would have given you anyway.

## What you get

- **Router**: Gin or Chi, a working `main.go` with a health check endpoint
- **Database**: Postgres, MySQL, MongoDB, or Redis, connection pooling code plus a docker-compose service and `.env.example`
- **Frontend**: Vite (React + TypeScript) or Next.js, scaffolded via `npm create vite@latest` or `create-next-app`, falling back to pnpm automatically if npm isn't on your machine
- **Docker**: an optional multi-stage Dockerfile for the backend, on request
- A `.gitignore` and `README.md` written for whatever you actually picked, not a generic template with every option commented out

Everything is opt-in. `--db=none --frontend=none` gets you a bare Go module with just a router. Nothing gets generated that you didn't ask for.

## Install

**macOS / Linux**
```bash
brew install sajanv88/tap/gocli
```

**Windows**
```powershell
scoop bucket add sajanv88 https://github.com/sajanv88/scoop-bucket
scoop install gocli
```

**Linux**
```bash
curl -LO https://github.com/sajanv88/gocli/releases/latest/download/gocli_0.1.0_linux_amd64.tar.gz
tar xzf gocli_0.1.0_linux_amd64.tar.gz
sudo mv gocli /usr/local/bin/
```


**Anywhere with Go installed**
```bash
go install github.com/sajanv88/gocli@latest
```

Or grab a binary directly from the [releases page](https://github.com/sajanv88/gocli/releases).

## Usage

Run it with no flags and it'll ask what it needs:

```bash
gocli new myapp
```

Or skip the prompts entirely:

```bash
gocli new myapp \
  --module=github.com/you/myapp \
  --router=chi \
  --db=postgres \
  --frontend=vite \
  --docker
```

| Flag             | Values                                          | Default |
|------------------|-------------------------------------------------|---|
| `--module`       | any Go module path                              | required |
| `--router`       | `gin`, `chi`, `none`                            | prompted |
| `--db`           | `postgres`, `mysql`, `mongodb`, `redis`, `none` | prompted |
| `--frontend`     | `vite`, `nextjs`, `none`                        | prompted |
| `--agent`        | `adk`, `eino`, `none`                           | prompted |
| `--docker`       | flag                                            | off |
| `--force`        | flag                                            | off, overwrites an existing directory |
| `--install-deps` | flag                                            | prompted if a frontend was generated |

Every flag you pass skips the matching prompt, so this runs fine in CI without ever hitting an interactive form.

{{if ne .Agent "none"}}
### AI agent ({{.Agent}})

```bash
cp agent/.env.example agent/.env
# add your Gemini API key from https://aistudio.google.com/app/apikey
```

```bash
source agent/.env   # or agent\env.bat on Windows
go run ./agent
```

Or with the interactive web UI:

```bash
go run ./agent web api webui
```
{{end}}


## What it's built on

gocli follows the same hexagonal architecture I use for actual production services: a domain layer that has no idea Cobra, os/exec, or embed.FS exist, a use case that orchestrates the whole scaffold, and one small adapter per router, database, and frontend option. Adding a fifth database means writing one new file and adding one line to a registry map. Nothing else in the codebase has to change.

## Ideas I'm chewing on

- A Dockerfile for the frontend too, not just the backend (Next.js needs `output: 'standalone'` set for a slim image, Vite needs an nginx stage)
- More routers (Fiber, Echo) and more databases if there's demand for them
- A `gocli list` command to print what's available without digging through `--help`
- Config file support, so a team can check in their preferred defaults instead of typing the same flags every time


None of these are promises, just what I'd reach for next. If one of them would actually solve a problem you have, open an issue and say so. That moves it up the list faster than me guessing.

## Contributing

If something breaks, or the generated code doesn't build, open an issue with the exact command you ran. If you want to add a router, database, or frontend option, the adapter pattern above means you're mostly writing templates, not touching core logic, so it's a reasonable first PR even if you haven't read the rest of the codebase.

Feedback on the CLI itself matters just as much as code. Tell me what's confusing about the prompts, or what flag you wished existed. I built this for my own workflow first, so there's a decent chance it's missing something obvious to someone using it differently.

## License

MIT. Use it, fork it, put your own name on the fork if that's more useful to you.