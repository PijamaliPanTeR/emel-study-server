# Emel Study Server

Backend for the Emel Study (Sound Map) academic project. Built with Go and Fiber, following the same structure as main-service (modules, pkg, conf).

## Structure

- `pkg/` – config, version, utils
- `conf/` – YAML config (e.g. `local.yml`)
- `modules/server_module/` – health check
- `modules/study_module/` – study API (sounds, session, map, answers)

## Setup

```bash
go mod tidy
```

## Run

From repo root:

```bash
go run application.go
```

Listens on `127.0.0.1:6100` by default (see `conf/local.yml`).

## API

- `GET /` – health check
- `GET /study/sounds` – list of 22 sounds (id, audioUrl, order)
- `POST /study/session` – create session, returns `{ "sessionId": "..." }`
- `POST /study/session/:id/map` – body `{ "positions": [{ "soundId", "x", "y" }] }`
- `POST /study/session/:id/answers` – body `{ "groupStrategy", "groupsRepresent" }`

Sessions and data are stored in memory. For production, add persistence (e.g. DB) in `study_service`.

## Sounds

The API returns sound entries with `audioUrl: "/sounds/s1.mp3"` … `/sounds/s22.mp3`. The client loads these from its own origin (`public/sounds/`). To serve sounds from this server, add a static file route for `/sounds/*` and optionally return full URLs in the sounds list.
