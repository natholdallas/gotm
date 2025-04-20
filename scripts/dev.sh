#! /usr/bin/bash

# choose a way you build application

go build -o ./bin/backend

# go watch support
# gowatch -o ./bin/backend

# tmux support
# tmux new-session -d -s dev -n server 'gowatch -o ./bin/backend'

# if you have website folder in your project
# tmux new-window -t dev:1 -n client 'cd web && pnpm dev'
