#!/bin/sh

set -x

# test
goapp get golang.org/x/tools/cmd/cover
goapp get github.com/stretchr/testify

# AppEngine

goapp get google.golang.org/appengine

# Other library

goapp get golang.org/x/net/context
goapp get -u github.com/unrolled/render
goapp get -u github.com/kyokomi/goslash/goslash
goapp get -u github.com/kyokomi/goslash/plugins/time
goapp get -u github.com/kyokomi/goslash/plugins/echo
goapp get -u github.com/kyokomi/goslash/plugins/suddendeath
goapp get -u github.com/kyokomi/goslash/plugins/lgtm
goapp get -u github.com/PuerkitoBio/goquery

