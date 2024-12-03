# Day 2

## Part One

This time I tried using the language from the puzzle prompt to guide the assistant. I was able to get the correct answer with two messages.

I modified the prompt slightly to remove the fun names from the puzzle like "The Red-Nosed reactor safety systems" and just called them "the report" instead, because I assumed those names would throw off the assistant. I might try leaving them in next time to see if it makes a difference.

Chat history:

```
1. User: Load the puzzle input from the file named "input". The puzzle input consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

This example data contains six reports each containing five levels.

2. Assistant: I'll help you write a program to analyze the reports. Here's a Go program that reads the reports from a file and checks if each one is safe according to the rules:

3. User: A report only counts as safe if both of the following are true:

    The levels are either all increasing or all decreasing.
    Any two adjacent levels differ by at least one and at most three.

In the example above, the reports can be found safe or unsafe by checking those rules:

    7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
    1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
    9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
    1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
    8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
    1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

So, in this example, 2 reports are safe.

Analyze the reports. How many reports are safe?

4. Assistant: Generated safe report checking code.

5. Human: Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the first example's reports are now safe:

    7 6 4 2 1: Safe without removing any level.
    1 2 7 8 9: Unsafe regardless of which level is removed.
    9 7 6 2 1: Unsafe regardless of which level is removed.
    1 3 2 4 5: Safe by removing the second level, 3.
    8 6 4 4 1: Safe by removing the third level, 4.
    1 3 6 7 9: Safe without removing any level.
handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?

Assistant: Added an additional step to check for unsafe rules.

```
    