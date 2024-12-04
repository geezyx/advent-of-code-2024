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
    - Example prompts:
      ```
      I have added part one of the dayXX puzzle to puzzle.md, please add test input and expected result from part 1, and then solve part 1.

      I have added part two of the dayXX puzzle to puzzle.md, please add test input and expected result from part 2, and then solve part 2.

      Please update prompts.md with a summary of this session and the solution.
      ```

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
