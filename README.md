# Advent of Code 2024 - AI edition

I'm using this year's Advent of Code challenges to practice AI prompt engineering skills with Cursor.

## Project Structure

Each `dayXX` folder contains:
- `puzzle.md` - Puzzle description
- `input` - Puzzle input
- `solution.go` - Solution implementation
- `solution_test.go` - Tests
- `prompts.md` - Chat history and observations

## Getting Started

1. Create a new day's files:
   ```sh
   make new-day N=1
   ```

2. Add puzzle description to `puzzle.md` and input to `input`

3. Use Cursor to help implement the solution in `solution.go`

4. Run tests and solution:
   ```sh
   make test N=1  # Run tests for day 1
   make run N=1   # Run solution for day 1
   ```

## Development Tools

The project includes:
- [Makefile](Makefile) with common commands for Windows/Linux compatibility
- [newday.go](scripts/newday.go) script to generate boilerplate files
- Test framework for validating solutions
- [Utility functions](pkg/utils/file.go) for file handling

For all available commands, run:
   ```sh
   make help
   ```

The key functionality is implemented in:
- [Makefile](Makefile) - Cross-platform build commands
- [scripts/newday.go](scripts/newday.go) - New day generator
