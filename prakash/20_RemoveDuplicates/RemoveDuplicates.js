function removeDuplicates(nums) {
  if (nums.length === 0) {
    return 0;
  }
  var j = 0; // This will keep track of the last unique element's index
  for (var i = 1; i < nums.length; i++) {
    if (nums[i] !== nums[j]) {
      j++; // Move `j` to the next position
      nums[j] = nums[i]; // Place the unique element at position `j`
    }
  }
  return j + 1; // Length of the array with unique elements
}

module.exports = removeDuplicates;
