#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
#CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && ./scripts/dinggo )"
tmux bind-key T run-shell "$CURRENT_DIR/scripts/dinggo_runner.sh"
