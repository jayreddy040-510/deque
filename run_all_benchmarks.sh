#!/usr/bin/env bash

run_benchmark () {
    package_path=$1
    echo "running benchmark for package $package_path"
    echo "running benchmark for package $package_path" >> benchmarks.txt
    go test "$package_path"  -bench . -benchmem >> benchmarks.txt
    echo "" >> benchmarks.txt
}

run_benchmark "github.com/jayreddy040-510/deque/dll"
run_benchmark "github.com/jayreddy040-510/deque/dllpool"
run_benchmark "github.com/jayreddy040-510/deque/circular"

echo "all benchmarks completed" >> benchmarks.txt
