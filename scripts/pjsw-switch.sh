#!/usr/bin/env bash

pjsw-project() {
    local dir
    if [[ -e "$1" ]]; then
        echo "switching using pjsw sw"
        dir=$(pjsw sw "$1")
    else
        dir=$(pjsw project)
    fi

    cd "$dir" || return
}
