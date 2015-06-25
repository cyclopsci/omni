#!/bin/bash
work=.cover

rm -rf "$work"
mkdir "$work"
for pkg in $(go list ./...); do
  f="$work/$(echo $pkg | tr / -).cover"
  go test -covermode="count" -coverprofile="$f" "$pkg"
done

echo "mode: count" >"$work/coverage.out"
grep -h -v "^mode:" "$work"/*.cover >>"$work/coverage.out"

go tool cover -html="$work/coverage.out"
