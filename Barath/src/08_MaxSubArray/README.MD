# Kadane’s Algorithm

Kadane's Algorithm is to keep track of the maximum sum of the subarray as we move through the array.

- Start from the first element.
- For each element, decide whether to add it to the previous subarray or start a new subarray from that element.
- Keep track of the maximum sum encountered at each step.

## How it works

1. Initialize two variables:
   - `maxSum` to keep track of the maximum sum we've found.
   - `currentSum` to keep track of the current subarray sum.

2. Iterate through the array:
   - For each element, decide whether to add it to the current sum (i.e., extend the subarray) or start a new subarray.
   - Update `maxSum` if `currentSum` is larger.

3. Return `maxSum` as the result.

