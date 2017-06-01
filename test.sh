#!/bin/bash

PASS=0
FAIL=0

function eq() {
    actual=$1
    expected=$2

    if [ "$actual" = "$expected" ]
    then
        PASS=$((PASS + 1))
    else
        echo "fail: actual: ${actual}    expected: ${expected}"
        FAIL=$((FAIL + 1))
    fi
}

function report() {
    echo "--------------------------------"
    echo "pass: ${PASS}    fail: ${FAIL}"
}

# tests
eq $(go run main.go + 1 2) 3
eq $(go run main.go - 1 2) -1
eq $(go run main.go \* 2 3) 6
eq $(go run main.go / 4 2) 2
eq $(go run main.go + 2 \* 3 5) 17
eq $(go run main.go \* 2 + 3 5) 16
report
