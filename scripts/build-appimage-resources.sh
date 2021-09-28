#!/bin/sh
# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -ex

: ${GO:=go}

binary=aaaaxy-$($GO env GOOS)-$($GO env GOARCH)

convert assets/sprites/riser_small_up_0.png \
	-filter Point -geometry 128x128 \
	packaging/"$binary.png"
sh scripts/aaaaxy.desktop.sh > packaging/"$binary.desktop"
sh scripts/io.github.divverent.aaaaxy.metainfo.xml.sh > packaging/io.github.divverent."$binary".metainfo.xml
