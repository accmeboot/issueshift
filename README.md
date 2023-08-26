# Issueshift
Agile board web app

## Development 

### Start dev server
```console
make run/dev
```

### Create migrations
```console
make migrations/create
```

### Run migrations
```console
make migrations/up
```
```console
make migrations/down
```

## Stack
- https://github.com/go-chi/chi
- https://github.com/pressly/goose
- https://github.com/sqlc-dev/sqlc

## Todo:
- [ ] Overall api architecture
- [ ] Stateful authorization
- [ ] Projects (workspaces?)
- [ ] Tasks
- [ ] Boards
- [ ] Account management
- [ ] Frontend
