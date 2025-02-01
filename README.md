# Scan Tool

## Overview

A simple command-line tool to scan a specified directory for files with a given extension and organize results.

## Usage

* Create a result directory (if not exists):
    - mkdir -p result

* Run the scan command:
    - ./scan {/path/of/directory} {extension}

## Parameters:

- /path/of/directory : The absolute or relative path of the directory to scan.

- extension : The file extension to filter (e.g., php, js).

## Example

- ./scan /home/user/documents php

This command scans /home/user/documents for .txt files and stores the results in the result directory.

* Requirements: Unix-based OS (Linux/macOS)

* Ensure the scan script has execution permissions: chmod +x scan

## License

This project is licensed under the Ninja Sawit.
