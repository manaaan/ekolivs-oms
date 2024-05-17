#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "Updating custom product types"

# File to be updated
file_path="zettle/zettle-product.gen.go"

# Ensure the file exists
if [[ ! -f $file_path ]]; then
    echo "Error: File $file_path not found!"
    exit 1
fi

# Ensure go is installed
if ! command -v go &> /dev/null; then
    echo "Error: go is not installed. Please install go to continue."
    exit 1
fi

# VatPercentage from *float32 to *string
sed -i '' -E 's/(VatPercentage[[:space:]]+)\*float32/\1*string/g' $file_path

# Updated from time.Time to *string
sed -i '' -E 's/(Updated[[:space:]]+)time.Time/\1*string/g' $file_path

# Created from time.Time to *string
sed -i '' -E 's/(Created[[:space:]]+)time.Time/\1*string/g' $file_path

# Remove leftover time import
sed -i '' -E 's/"time"//g' $file_path

# Format the file
go fmt $file_path

echo "Done updating custom product types"