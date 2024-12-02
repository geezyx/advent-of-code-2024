

# Day 1

## Part 1

First I put the puzzle input into a file named `input`. Then I created a main.go file and began chatting with the assistant to complete the puzzle.

Chat history:
```
1. User: Write a simple main function that loads the contents of a file named "input" into a slice of bytes.
2. Assistant: Provided code to read file into a slice of bytes.
3. User: Parse the data into two slices of integers.
4. Assistant: Provided code to parse data into two slices of integers.
5. User: Order the two slices from smallest to largest.
6. Assistant: Provided code to sort the slices.
7. User: Iterate over both slices and sum the distances.
8. Assistant: Provided code to calculate and sum distances.
```

### Notes

Up until the final message to the assistant, I used the full generated code to continue building the solution.

After my final message to the assistant, the assistant changed the code to use example slices instead of the puzzle input. I had to combine the code I had already generated with the new code to complete the puzzle.

## Part 2

I was able to complete this part with a single message to the assistant. I just slightly modified the question from the puzzle prompt to match the variable names used in the code.

Chat history:
```
1. User: Calculate a total similarity score by adding up each number in firstInts after multiplying it by the number of times that number appears in the secondInts.
2. Assistant: Provided code to calculate the total similarity score.
```
