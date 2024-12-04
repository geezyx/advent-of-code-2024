# Day 3: Mull It Over

## Part One

Chat history:
```
1. User: Shared puzzle description and asked to add test input and solve part 1
2. Assistant: Added test case with example input "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))" and implemented solution using regex
3. User: Accepted the solution which correctly identified valid mul instructions and calculated sum of 161
```

Key observations:
- Regular expression pattern `mul\((\d{1,3}),(\d{1,3})\)` effectively filtered valid instructions
- Test-driven development with example data
- Solution handled parsing and multiplication in a clean, error-checked way

Implementation reference:
startLine: 29
endLine: 61

## Part Two

Chat history:
```
1. User: Added part two requirements for handling do() and don't() instructions
2. Assistant: Extended solution with instruction type and state tracking
3. User: Accepted the solution which correctly calculated sum of 48 for enabled multiplications
```

Key observations:
- Added instruction struct to track type and position
- Combined regex pattern to match all instruction types
- Maintained enabled/disabled state while processing instructions in order
- Reused multiplication parsing logic from part 1

Implementation reference:
startLine: 63
endLine: 105

## Solution Evolution
Initial approach:
- Started with regex pattern for valid mul instructions
- Added structured type for instruction parsing
- Implemented state tracking for part 2

## Learnings
- What worked well:
  - Regular expressions for parsing corrupted input
  - Clear instruction type definitions
  - State tracking for conditional execution
- Ideas for future improvements:
  - Could potentially combine regex patterns to avoid repeated compilation
  - Might benefit from more test cases with edge cases

## Time Analysis
- Time spent on prompt engineering: ~3 minutes
- Time spent on implementation: ~10 minutes
- Test-driven approach helped ensure correct implementation

## Code Stats
- Lines of code: ~105 lines
- Number of prompts needed: 2 total (1 for each part)
- Number of iterations to solution: 1 for part 1, 1 for part 2
