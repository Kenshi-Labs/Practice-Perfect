export default function rotate(nums: any, k: any) {
  const n = nums.length;
  if (n === 0) return; // Edge case for an empty array

  k = k % n; // Handle rotation larger than the array length

  // Helper function to reverse part of the array
  function reverse(start: any, end: any) {
    while (start < end) {
      [nums[start], nums[end]] = [nums[end], nums[start]]; // Swap
      start++;
      end--;
    }
  }

  // Step 1: Reverse the entire array
  reverse(0, n - 1);

  // Step 2: Reverse the first k elements
  reverse(0, k - 1);

  // Step 3: Reverse the remaining elements
  reverse(k, n - 1);
}
