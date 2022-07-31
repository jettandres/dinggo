#!/usr/bin/env bash

get_tmux_option() {
    local option=$1
    local default_value=$2
    local option_value="$(tmux show-option -gqv "$option")"

    if [[ -z "$option_value" ]]; then
        echo "$default_value"
    else
        echo "$option_value"
    fi
}

do_interpolation() {
    echo "hatdog"
}

update_tmux_option() {
	local option=$1
	local option_value=$(get_tmux_option "$option")
	local new_option_value=$(do_interpolation "$option_value")
        tmux set-option -gq "github_notif" "1 test notif"
}

main() {
	update_tmux_option "status-right"
	update_tmux_option "status-left"
}

main
