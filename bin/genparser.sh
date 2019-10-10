#!/usr/bin/env bash

CWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source "${CWD}/common.sh"

${CWD}/antlr.sh -o ${ROOT}/gen/parser -Dlanguage=Go -listener -visitor -lib ${ROOT} ${ROOT}/Nebularis.g4