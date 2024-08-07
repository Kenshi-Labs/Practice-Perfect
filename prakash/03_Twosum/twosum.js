// const nums = [2, 7, 11, 15];
// const target = 9;
// const result = twosum(nums, target);
// const result2 = twosum(nums, target);
// console.log(result); 1st case
// console.log(result2); // 2nd case

function twosum(nums, target) {
  // 1st case
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] + nums[i + 1] === target) {
      return [i, i + 1];
    }
  }

  // 2nd Case
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] + nums[j] === target) {
        return [i, j];
      }
    }
  }
  return [];
}

module.exports = twosum;
