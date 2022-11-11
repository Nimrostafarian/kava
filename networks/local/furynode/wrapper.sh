#!/usr/bin/env sh

BINARY=/fury/linux/${BINARY:-furyd}
echo "binary: ${BINARY}"
ID=${ID:-0}
LOG=${LOG:-fury.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'kvd' E.g.: -e BINARY=kvd_my_test_version"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export FURYHOME="/fury/node${ID}/fury"

if [ -d "$(dirname "${FURYHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${FURYHOME}" "$@" | tee "${FURYHOME}/${LOG}"
else
  "${BINARY}" --home "${FURYHOME}" "$@"
fi
