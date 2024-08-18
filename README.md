# EveCal README

## Scripts

`lsof -i tcp:3000` find current process on that port

## Goose Commands

all using `./goose.sh`

`create init sql` change init with a name

`status` Get current migration status

`up` Apply all available migrations.

`down` Roll back a single migration from the current version.

`redo` Roll back the most recently applied migration, then run it again.

## Develop

 Choose a Task command to run

  build             build a binary
  dev               build and run local project
  container         build a container
  sqlc              generate sqlc files
