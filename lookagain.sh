#!/bin/bash
find . -type f -name "*.sh" | cut -d "." -f2 | rev | cut -d "/" -f1 | rev | sort -r
