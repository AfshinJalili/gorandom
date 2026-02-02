#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
HOOKS_DIR="${ROOT_DIR}/.githooks"

if [ ! -d "${HOOKS_DIR}" ]; then
  echo "Missing ${HOOKS_DIR}. Did you clone the repo correctly?"
  exit 1
fi

git config core.hooksPath "${HOOKS_DIR}"
chmod +x "${HOOKS_DIR}/pre-commit"

echo "Git hooks installed from ${HOOKS_DIR}"
