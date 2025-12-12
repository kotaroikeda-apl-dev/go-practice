#!/bin/bash

# 全ての例を実行するスクリプト

examples=(
	"struct"
	"function"
	"nonstruct"
	"pointer"
	"indirection"
	"indirection-value"
	"receiver-choice"
)

for example in "${examples[@]}"; do
	echo "Running $example..."
	cd "$example" && go run main.go
	cd ..
	echo ""
done

