#!/usr/bin/env sh
set -o errexit
set -o nounset

git config commit.template .gitmessage

cp "${PROJECT_PATH}/scripts/pre-commit" "${PROJECT_PATH}/.git/hooks/"
