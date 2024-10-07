function ArrayIntersection(nums1, nums2, nums3) {
  var set1 = new Set(nums1);
  var resultSet = new Set();
  for (var _i = 0, nums2_1 = nums2; _i < nums2_1.length; _i++) {
    var num = nums2_1[_i];
    if (set1.has(num)) {
      resultSet.add(num);
    }
  }
  return Array.from(resultSet);
}

module.exports = ArrayIntersection;
