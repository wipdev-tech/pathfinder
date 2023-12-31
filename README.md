```
    ____          __   __     ____ _             __           
   / __ \ ____ _ / /_ / /_   / __/(_)____   ____/ /___   _____
  / /_/ // __ `// __// __ \ / /_ / // __ \ / __  // _ \ / ___/
 / ____// /_/ // /_ / / / // __// // / / // /_/ //  __// /    
/_/     \__,_/ \__//_/ /_//_/  /_//_/ /_/ \__,_/ \___//_/     
```                                                                


A simple CLI for managing path variables

## Installation

### Binary

*Coming Soon™*

### From source

- Ensure Go is installed
- `cd` into the project root directory and run

```bash
go install ./cmd/pathfinder/
```

## Usage

> **Disclaimer:**
> As of right now, this tool has been only tested on Linux.

- To list available commands run the following

```bash
pathfinder help  # `pathfinder` only will also work
```

- To list paths in your PATH variable run the following

```bash
pathfinder list
```

- To *temporarily* add a new path to your PATH variable run the following

```bash
pathfinder add /your/new/path
```

- To *permanently* add a new path to your PATH variable run the following
    + This adds a command to your `.bashrc`, and you will be prompted before proceeding
    + The PATH variable is also modified for the current session - no restart needed

```bash
pathfinder add /your/new/path --persist
```

