function maxSubArray(nums: number[]): number {
  if (nums.length === 0)
    throw new Error("Array must contain at least one number");

  let maxSoFar = nums[0];
  let currentMax = nums[0];

  for (let i = 1; i < nums.length; i++) {
    currentMax = Math.max(nums[i], currentMax + nums[i]);
    maxSoFar = Math.max(maxSoFar, currentMax);
  }

  return maxSoFar;
}

export default maxSubArray;
