#!/bin/bash
VERBOSE=${VERBOSE:-"0"}
V=""
if [[ "${VERBOSE}" == "1" ]];then
    V="-x"
    set -x
fi

SCRIPTPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

APP_NAME=${1:?"app name"}
OUT=${2:?"output path"}

shift
set -e
# GOOS=${GOOS:-darwin}
# GOARCH=${GOARCH:-amd64}
GOBINARY=${GOBINARY:-go}
# GOPKG="$(go env GOPATH)/pkg"
BUILDINFO=${BUILDINFO:-""}
STATIC=${STATIC:-1}
LDFLAGS="-s -w -extldflags -static"
GOBUILDFLAGS=${GOBUILDFLAGS:-""}
GCFLAGS=${GCFLAGS:-}

# Split GOBUILDFLAGS by spaces into an array called GOBUILDFLAGS_ARRAY.
IFS=' ' read -r -a GOBUILDFLAGS_ARRAY <<< "$GOBUILDFLAGS"

GCFLAGS=${GCFLAGS:-}

if [[ "${STATIC}" !=  "1" ]];then
    LDFLAGS="-s -w"
fi

# gather buildinfo if not already provided
# For a release build BUILDINFO should be produced
# at the beginning of the build and used throughout
if [[ -z ${BUILDINFO} ]];then
    BUILDINFO=$(mktemp)
    "${SCRIPTPATH}/report_build_info.sh"  ${APP_NAME}> "${BUILDINFO}"
fi

# BUILD LD_EXTRAFLAGS
LD_EXTRAFLAGS=""

while read -r line; do
    LD_EXTRAFLAGS="${LD_EXTRAFLAGS} -X ${line}"
done < "${BUILDINFO}"

# verify go version before build
# NB. this was copied verbatim from Kubernetes hack
minimum_go_version=go1.13 # supported patterns: go1.x, go1.x.x (x should be a number)
IFS=" " read -ra go_version <<< "$(${GOBINARY} version)"
if [[ "${minimum_go_version}" != $(echo -e "${minimum_go_version}\n${go_version[2]}" | sort -s -t. -k 1,1 -k 2,2n -k 3,3n | head -n1) && "${go_version[2]}" != "devel" ]]; then
    echo "Warning: Detected that you are using an older version of the Go compiler. APP requires ${minimum_go_version} or greater."
fi

OPTIMIZATION_FLAGS="-trimpath"
if [ "${DEBUG}" == "1" ]; then
    OPTIMIZATION_FLAGS=""
fi

#goxc -d=${OUT} -pv=${APP_VERSION} -bc='linux,windows,darwin' -build-ldflags="${LD_EXTRAFLAGS}"
echo time ${GOBINARY} build ${OPTIMIZATION_FLAGS} ${V} ${GOBUILDFLAGS} ${GCFLAGS:+-gcflags "${GCFLAGS}"} -o ${OUT} \
       -ldflags "${LDFLAGS} ${LD_EXTRAFLAGS}"
time ${GOBINARY} build ${OPTIMIZATION_FLAGS} ${V} ${GOBUILDFLAGS} ${GCFLAGS:+-gcflags "${GCFLAGS}"} -o ${OUT} \
       -ldflags "${LDFLAGS} ${LD_EXTRAFLAGS}"