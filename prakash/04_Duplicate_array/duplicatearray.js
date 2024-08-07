function duplicatearray(nums) {
  if (nums.length === 0) return 0;
  var uniqueIndex = 1;
  for (var i = 1; i < nums.length; i++) {
    if (nums[i] !== nums[i - 1]) {
      nums[uniqueIndex] = nums[i];
      uniqueIndex++;
    }
  }
  return uniqueIndex;
}

module.exports = duplicatearray;
