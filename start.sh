#!/bin/bash

cd /location-calculator || { echo "Directory not found"; exit 1; }

# Run the Go file
go run location-calculator.go