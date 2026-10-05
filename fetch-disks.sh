#!/usr/bin/env bash
#
# Downloads the disks the activities use into disks/, as listed in disks.tsv.
#
# It can be run any number of times: a disk already in disks/ with the
# checksum of the list is left alone, and only what is missing or has changed
# is downloaded. A disk with "-" as its checksum is downloaded when missing,
# and its checksum printed to be put in the list.
#
#   ./fetch-disks.sh            all the disks
#   ./fetch-disks.sh blank.dsk  only the ones named

set -euo pipefail

cd "$(dirname "$0")"
list=disks.tsv
target=disks
mkdir -p "$target"

if command -v sha256sum >/dev/null; then
    checksum() { sha256sum "$1" | cut -d' ' -f1; }
else
    checksum() { shasum -a 256 "$1" | cut -d' ' -f1; }
fi

# wanted says whether a disk was named on the command line
wanted() {
    local name=$1
    for w in "${selected[@]}"; do
        [ "$w" = "$name" ] && return 0
    done
    return 1
}

selected=("$@")
failed=0

while IFS=$'\t' read -r name sum source member; do
    case "$name" in ''|'#'*) continue ;; esac
    if [ ${#selected[@]} -gt 0 ] && ! wanted "$name"; then
        continue
    fi

    file="$target/$name"
    if [ -f "$file" ]; then
        if [ "$sum" = "-" ] || [ "$(checksum "$file")" = "$sum" ]; then
            echo "ok        $name"
            continue
        fi
        echo "changed   $name"
    else
        echo "fetching  $name"
    fi

    work=$(mktemp -d)
    case "$source" in
    blank:*)
        head -c "${source#blank:}" /dev/zero >"$work/disk"
        ;;
    *)
        if ! curl -fsSL --retry 3 -o "$work/download" "$source"; then
            echo "  could not download $source" >&2
            rm -rf "$work"; failed=1; continue
        fi
        if [ -n "${member:-}" ]; then
            if ! unzip -p "$work/download" "$member" >"$work/disk"; then
                echo "  $member is not in $source" >&2
                rm -rf "$work"; failed=1; continue
            fi
        else
            mv "$work/download" "$work/disk"
        fi
        ;;
    esac

    got=$(checksum "$work/disk")
    if [ "$sum" != "-" ] && [ "$got" != "$sum" ]; then
        echo "  wrong checksum for $name: $got, expected $sum" >&2
        rm -rf "$work"; failed=1; continue
    fi
    if [ "$sum" = "-" ]; then
        echo "  sha256 $got"
    fi
    mv "$work/disk" "$file"
    rm -rf "$work"
done <"$list"

exit $failed
