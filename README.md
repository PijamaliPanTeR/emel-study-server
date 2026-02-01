# Emel Study Server

Backend for the Emel Study (Sound Map) academic project. Built with Go and Fiber (modules, pkg, conf).

## Structure

- `pkg/` – config, version, utils
- `conf/` – YAML config (e.g. `local.yml`)
- `modules/server_module/` – health check
- `modules/study_module/` – study API (session, map, answers, progress)

## Setup

```bash
go mod tidy
```

## Run

From repo root:

```bash
go run application.go
```

Listens on `127.0.0.1:7100` by default (see `conf/local.yml`).

## API

- `GET /` – health check
- `POST /study/session` – create session, returns `{ "sessionId": "..." }` (optional progress for resume)
- `POST /study/session/:id/map` – body `{ "positions": [{ "soundId", "x", "y" }] }`
- `POST /study/session/:id/answers` – body `{ "groupStrategy", "groupsRepresent" }`
- `POST /study/session/:id/progress` – body `{ "currentStep", "listenedSoundIds", "soundGroups", "defineGroupsRectangles" }` (all optional, merge)
