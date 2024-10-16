function removeDuplicates(nums: (number | string)[]): number {
  if (nums.length === 0) return 0;

  let i = 0;
  for (let j = 1; j < nums.length; j++) {
    if (nums[j] !== nums[i]) {
      i++;
      nums[i] = nums[j];
    }
  }

  // Fill the rest of the array with "_"
  for (let k = i + 1; k < nums.length; k++) {
    nums[k] = "_";
  }

  return i + 1; // New length of the array with unique elements
}

export default removeDuplicates;
