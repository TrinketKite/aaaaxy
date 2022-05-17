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

if git ls-files -dmo | grep -q .; then
	echo >&2 'Working directory is not clean. Please commit or clean first.'
	exit 1
fi

prev=$(git describe --always --long --match 'v*.*' --exclude 'v[0-9].[0-9]' --exclude 'v[0-9].[0-9].0-alpha' --exclude 'v[0-9].[0-9].0-beta' --exclude 'v[0-9].[0-9].0-rc')
# We want to exclude v*.* and v*.*.0-(alpha/beta).
prev=${prev%-*-g*}

new=$(sh scripts/version.sh gittag)

cat <<EOF >.commitmsg
Release $new

Changes since $prev:
$(git log --format='%w(72,2,4)- %s' "$prev"..)
EOF
vi .commitmsg

# Update gamecontroller mappings.
git submodule update --remote

# Include exact versions of submodules so that the source tarball on github
# contains the exact submodule version info.
git submodule > .gitmoduleversions

# Also store the current semver in the checkout. Used for compiling from
# source tarballs.
sh scripts/version.sh semver > .lastreleaseversion

# Update metainfo with current date and version already, and replace the text by a placeholder.
VERSION=$new DATE=$(date +%Y-%m-%d) MSG=$(cat .commitmsg) perl -0777 -pi -e '
	use strict;
	use warnings;
	my $version = $ENV{VERSION};
	my $date = $ENV{DATE};
	my $msg = $ENV{MSG};
	$msg =~ s/^Release .*//gm;
	$msg =~ s/^Changes since .*//gm;
	$msg =~ s/^  - /<\/li><li>/gm;
	$msg =~ s/^    //gm;
	$msg =~ s/^\n*<\/li>/<ul>/s;
	$msg =~ s/\n*$/<\/li><\/ul>/s;
	$msg =~ s/\n*<\/li>/<\/li>/g;
	$msg =~ s/\n/ /g;
	s/releases\/[^\/<]*<\/url>/releases\/$version<\/url>/g;
	s/<release version="[^"]*" date="[0-9-]*">/<release version="$version" date="$date">/g;
	s/<description>.*<\/description>/<description>$msg<\/description>/g;
' io.github.divverent.aaaaxy.metainfo.xml

# Also pack the SDL game controller DB at the exact version used for the
# release. Used for compiling from source tarballs.
7za a -tzip -mx=9 sdl-gamecontrollerdb-for-aaaaxy-$new.zip third_party/SDL_GameControllerDB/assets/input/*

# Also pack the files that do NOT get embedded into a mapping pack.
(
	cd assets/
	7za a -tzip -mx=9 ../mappingsupport-for-aaaaxy-$new.zip ../LICENSE objecttypes.xml _* */_*
)

GOOS=linux scripts/binary-release-compile.sh amd64
GOOS=windows scripts/binary-release-compile.sh amd64
GOOS=windows scripts/binary-release-compile.sh 386
# Note: sync the MACOSX_DEPLOYMENT_TARGET with current Go requirements and Info.plist.sh.
# Note: reduce deployment target back to 10.13 when https://github.com/hajimehoshi/ebiten/issues/2095 and/or https://github.com/tpoechtrager/osxcross/issues/346 are fixed.
GOOS=darwin CGO_ENV_amd64="PATH=$HOME/src/osxcross-sdk/bin:$PATH CGO_ENABLED=1 CC=o64-clang CXX=o64-clang++ MACOSX_DEPLOYMENT_TARGET=10.14" CGO_ENV_arm64="PATH=$HOME/src/osxcross-sdk/bin:$PATH CGO_ENABLED=1 CC=oa64-clang CXX=oa64-clang++ MACOSX_DEPLOYMENT_TARGET=10.14" LIPO="$HOME/src/osxcross-sdk/bin/lipo" scripts/binary-release-compile.sh amd64 arm64
GOOS=js scripts/binary-release-compile.sh wasm

git commit -a -m "$(cat .commitmsg)"
git tag -a "$new" -m "$(cat .commitmsg)"
newrev=$(git rev-parse HEAD)

git push origin tag "$new"

set +x

cat <<EOF
Please wait for automated tests on
https://github.com/divVerent/aaaaxy/actions

If these all pass, proceed by running

  scripts/publish-release.sh $new $newrev
EOF
