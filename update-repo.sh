#!/bin/bash

# Check if the script received exactly one argument
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 new_repo_location"
  exit 1
fi

# The new repository location provided as the argument
new_repo_location=$1

# The old repository URL to be replaced
old_repo_url="github.com/dlukt/graphql-backend-starter"

# Use grep to find all files containing the old repository URL
# and use xargs to pass the file names to sed for replacement
grep -rl "$old_repo_url" . | xargs sed -i "s|$old_repo_url|$new_repo_location|g"

echo "Replaced all occurrences of '$old_repo_url' with '$new_repo_location'."