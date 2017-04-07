#!/bin/bash

set +v

compilecmp -n 5
compilecmp -n 50 -cpu
compilecmp -n 30 -cpu -all -run=StdCmd

compilecmp -n 5 -beforeflags="-c=1" -afterflags="-c=2" HEAD HEAD
compilecmp -n 5 -beforeflags="-c=1" -afterflags="-c=4" HEAD HEAD
compilecmp -n 5 -beforeflags="-c=1" -afterflags="-c=8" HEAD HEAD
compilecmp -n 5 -beforeflags="-c=1" -afterflags="-c=16" HEAD HEAD

compilecmp -n 50 -cpu -beforeflags="-c=1" -afterflags="-c=2" HEAD HEAD
compilecmp -n 50 -cpu -beforeflags="-c=1" -afterflags="-c=4" HEAD HEAD
compilecmp -n 50 -cpu -beforeflags="-c=1" -afterflags="-c=8" HEAD HEAD
compilecmp -n 50 -cpu -beforeflags="-c=1" -afterflags="-c=16" HEAD HEAD
compilecmp -n 30 -cpu -all -run=StdCmd -beforeflags="-c=1" -afterflags="-c=2" HEAD HEAD
compilecmp -n 30 -cpu -all -run=StdCmd -beforeflags="-c=1" -afterflags="-c=4" HEAD HEAD
compilecmp -n 30 -cpu -all -run=StdCmd -beforeflags="-c=1" -afterflags="-c=8" HEAD HEAD
compilecmp -n 30 -cpu -all -run=StdCmd -beforeflags="-c=1" -afterflags="-c=16" HEAD HEAD
