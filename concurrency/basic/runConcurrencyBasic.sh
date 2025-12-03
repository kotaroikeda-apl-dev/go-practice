#!/bin/bash

# 全ての例を実行するスクリプト

examples=(
	"goroutine"
	"channelBasic"
	"channelBuffered"
	"channelRangeClose"
	"channelSelect"
	"mutex"
)

for example in "${examples[@]}"; do
	echo "Running $example..."
	cd "$example" && go run main.go
	cd ..
	echo ""
done

