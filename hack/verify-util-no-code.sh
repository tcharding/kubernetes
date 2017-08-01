#!/bin/bash

# Copyright 2014 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

shopt -s nullglob dotglob     # Include hidden files

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
DIR="${KUBE_ROOT}/pkg/util"

ALLOWABLE_FILES=(
    "BUILD"
)

function is-allowed() {
    local file=$1

    for allowed in ${ALLOWABLE_FILES[@]}
    do
        if [[ $allowed == $file ]]; then
            return 0
        fi
    done

    return 1
}

cd $DIR

for file in $(ls)
do
    if [ -f "$file" ]; then
       is-allowed $file
       if [[ $? != 0 ]]; then
           exit 1
       fi
    fi
done

exit 0
