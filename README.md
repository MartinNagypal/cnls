# CNLS — ChickenNugget ls

A small command-line directory lister written in Go. CNLS displays directory entries in columns, with options for showing hidden files or listing file details.

## Features

- Lists non-hidden entries in a directory by default
- Sorts entries alphabetically, without case sensitivity
- Displays entries in columns based on the terminal width
- Supports showing hidden files
- Provides a long format with permissions, type ,owner, group, modification time, size, and filename

## Requirements

- Go

## Build

Clone the repository and build the executable:

```sh
cd cnls
go build .