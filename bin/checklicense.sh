#!/usr/bin/env bash

CWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
source "${CWD}/common.sh"

PKG_FILES=$(find  ${ROOT}/pkg -name '*.go')
TESTS_FILES=$(find  ${ROOT}/tests -name '*.go')
MODULE_FILES=$(find  ${ROOT}/modules -name '*.ns')

ret=0
for FILE in ${PKG_FILES} ${TESTS_FILES} ${MODULE_FILES}
do
  if ! grep -L -q -e "Apache License, Version 2" "${FILE}"; then
    echo "Missing license: ${FILE}"
    ret=1
  fi

  if ! grep -L -q -e "Copyright" "${FILE}"; then
    echo "Missing copyright: ${FILE}"
    ret=1
  fi
done

exit ${ret}
