#!/bin/bash

INITIAL_DIR=$(pwd)

TITLE_DIR="cool_ascii_title"
TITLE_SCRIPT="main.go"

PART1_DIR="get_new_swagger"
PART1_SCRIPT="main.go"

PART2_DIR="get_reference_objects"
PART2_SCRIPT="main.go"

PART3_DIR="prepare_json"
PART3_SCRIPT="main.go"

PART4_DIR="compare_swaggers_python"
PART4_SCRIPT="main.py"

PART5_DIR="update_anchor"
PART5_SCRIPT="main.go"

PART6_DIR="summary"
PART6_SCRIPT="main.go"

PART7_DIR="cleanup"
PART7_SCRIPT="main.go"



# Function to run a script
run_script() {
    local dir="$1"
    local script="$2"
    
    echo "Running $script"
    cd "$dir" || exit 1
    
    if [[ "$script" == *.py ]]; then
        python3 "$script"
    else
        go run "$script"
    fi
    
    if [ $? -ne 0 ]; then
        echo "Script failed."
        exit 1
    fi
    
    cd "$INITIAL_DIR" || exit 1
}

# Run the scripts sequentially
run_script "$TITLE_DIR" "$TITLE_SCRIPT"
run_script "$PART1_DIR" "$PART1_SCRIPT"
run_script "$PART2_DIR" "$PART2_SCRIPT"
run_script "$PART3_DIR" "$PART3_SCRIPT"
run_script "$PART4_DIR" "$PART4_SCRIPT"
run_script "$PART5_DIR" "$PART5_SCRIPT"
run_script "$PART6_DIR" "$PART6_SCRIPT"
run_script "$PART7_DIR" "$PART7_SCRIPT"

echo "All scripts completed successfully"
