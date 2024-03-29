## Scripts
`sqlc generate` Generates new database files

`lsof -i tcp:3000` find current process on that port

## Goose Commands
all using `./goose.sh`
`create init sql` change init with a name

`status` Get current migration status

`up` Apply all available migrations.

`down` Roll back a single migration from the current version.

`redo` Roll back the most recently applied migration, then run it again.

## Develop

```
 Choose a make command to run

  init          initialize project (make init module=github.com/user/project)
  vet           vet code
  test          run unit tests
  build         build a binary
  dockerbuild   build project into a docker container image
  start         build and run local project
  css           build tailwindcss
  css-watch     watch build tailwindcss
```
