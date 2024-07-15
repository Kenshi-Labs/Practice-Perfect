function twosum(nums: number[], target: number): number[] | null {
  // Initialize an empty map to store numbers and their indices
  const numMap: { [key: number]: number } = {};

  // Iterate over the array of numbers

  for (let i = 0; i < nums.length; i++) {
    // Calculate the complement of the current number

    const complement: number = target - nums[i];

    // Check if the complement exists in the map

    if (numMap.hasOwnProperty(complement)) {
      // hasOwnProperty return boolean has its checks if the complement already exists in numMap;

      // If it exists, return the indices of the two numbers

      return [numMap[complement], i];
    }

    // Store the current number and its index in the map

    numMap[nums[i]] = i;
  }

  // If no solution is found, throw Error
  throw new Error("No two sum solution");
}
