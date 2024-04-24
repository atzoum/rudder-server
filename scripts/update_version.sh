#!/bin/bash
sed -E "s,(version += \")[^\"]*(\"),\1${1}\2," main.go > main.go.tmp
mv main.go.tmp main.go