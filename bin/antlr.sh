#!/usr/bin/env bash

CWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source "${CWD}/common.sh"

CLASSPATH="${CWD}/../lib/antlr-4.7.1-complete.jar:$CLASSPATH"
ANTLR="java -Xmx500M -cp ${CLASSPATH} org.antlr.v4.Tool"

${ANTLR} $*