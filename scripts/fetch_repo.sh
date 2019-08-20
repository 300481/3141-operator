#!/usr/bin/env bash

set -euo pipefail

REPOSITORY=${1}
COMMIT=${2}
PUSH_TS=${3}

SYSTEM_ID=$(kubectl --namespace 3141-operator get configmap 3141-operator-environment -o jsonpath='{.data.SYSTEM_ID}')

TEMPDIR=$(mktemp -d)

cd ${TEMPDIR}

git clone ${REPOSITORY} repo

cd repo

git checkout ${COMMIT}

cd ${SYSTEM_ID}

cat components.yaml