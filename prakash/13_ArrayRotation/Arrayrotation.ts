function rotateArray(nums: number[], k: number): void {
  const n = nums.length;
  k = k % n; // If k is greater than array length, reduce it
  if (n === 0 || k === 0) return;

  // Helper function to reverse a portion of the array
  function reverse(start: number, end: number): void {
    while (start < end) {
      const temp = nums[start];
      nums[start] = nums[end];
      nums[end] = temp;
      start++;
      end--;
    }
  }

  // Step 1: Reverse the entire array
  reverse(0, n - 1);

  // Step 2: Reverse the first k elements
  reverse(0, k - 1);

  // Step 3: Reverse the remaining n - k elements
  reverse(k, n - 1);
}
