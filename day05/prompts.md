# Day 5

## Part 1

Chat history:
```
1. User: Shared puzzle description and requested test input and solution for part 1
2. Assistant: Created test file with example input and expected result (143)
3. User: Reported input parsing error
4. Assistant: Fixed newline handling in parseInput function
5. User: Accepted solution implementation
```

Key observations:
- Input has two sections separated by double newline: rules and updates
- Rules format: X|Y means X must come before Y
- Solution validates update order against applicable rules
- Only rules where both numbers exist in update are considered
- Middle number from valid updates are summed

Solution referenced in:
```
startLine: 28
endLine: 85
```

## Part 2

Chat history:
```
1. User: Added part 2 puzzle description
2. Assistant: Updated test file with part 2 expected result (123)
3. Assistant: Implemented topological sort solution
4. User: Accepted solution implementation
```

Key observations:
- Part 2 focuses on fixing incorrectly ordered updates
- Used topological sort (Kahn's algorithm) to reorder updates
- Only process updates that were invalid in part 1
- Solution handles example cases correctly:
  ```
  75,97,47,61,53 -> 97,75,47,61,53 (middle: 47)
  61,13,29 -> 61,29,13 (middle: 29)
  97,13,75,29,47 -> 97,75,47,29,13 (middle: 47)
  ```
- Sum of middle numbers from reordered updates: 47 + 29 + 47 = 123

Solution referenced in:
```
startLine: 87
endLine: 165
```

## Solution Evolution
- Initial implementation for Part 1 focused on parsing input and validating order.
- Encountered input parsing issues due to newline inconsistencies, which were resolved by normalizing newlines.
- Part 2 required reordering updates using topological sorting, which was implemented using Kahn's algorithm.

## Learnings
- Handling input parsing robustly is crucial, especially with varying newline formats.
- Topological sorting is effective for resolving dependency orderings in graph-like problems.
- Clear separation of parsing, validation, and processing logic aids in debugging and extending solutions.

## Time Analysis
### Implementation Time
- Part 1: ~15 minutes (including fixing input parsing issues)
- Part 2: ~10 minutes
- Total time: ~25 minutes

### Computational Complexity
- Part 1: O(n * r) where n is number of updates and r is number of rules
- Part 2: O(n * (v + e)) where v is number of pages and e is number of rules in each update

## Code Stats
- Lines of code for Part 1: 58
- Lines of code for Part 2: 78
- Total lines of code: 136
- Functions implemented: 3 (parseInput, isValidOrder, topologicalSort)
