#!/bin/sh

set -e

destdir="$1"; shift

root=$PWD

tmpdir=$(mktemp -d)
trap 'rm -rf "$tmpdir"' EXIT

logged() {
	printf >&2 '+ %s\n' "$*"
	"$@"
}

if $NO_VFS; then
	logged touch "$tmpdir/use-local-assets.stamp"
else
	for sourcedir in assets third_party/*/assets; do
		logged cd "$root/$sourcedir"
		find . -type f | while read -r file; do
			mkdir -p "$tmpdir/${file%/*}"
			logged ln -snf "$root/$sourcedir/$file" "$tmpdir/$file"
		done
		directory=${sourcedir%/*}
		directory=${directory##*/}
		{
			echo "Applying to the following files:"
			logged find . -type f -print
			echo
			if [ -f ../COPYRIGHT.md ]; then
				logged cat ../COPYRIGHT.md
				echo
			fi
			logged cat ../LICENSE
		} | logged dd status=none of="$tmpdir/$directory.COPYRIGHT"
	done
fi
logged statik -m -f -src "$tmpdir/" -dest "$root/$destdir"
