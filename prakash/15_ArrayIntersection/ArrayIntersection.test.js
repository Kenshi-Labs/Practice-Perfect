const arrayIntersection = require("./ArrayIntersection");

describe("Array Intersection", () => {
  test("Test with arrays having common elements", () => {
    const nums1 = [1, 2, 2, 1];
    const nums2 = [2, 2, 3];
    const result = arrayIntersection(nums1, nums2);
    expect(result).toEqual([2]);
  });

  test("Test with arrays having no common elements", () => {
    const nums1 = [1, 2, 3];
    const nums2 = [4, 5, 6];
    const result = arrayIntersection(nums1, nums2);
    expect(result).toEqual([]);
  });

  test("Test with empty arrays", () => {
    const nums1 = [];
    const nums2 = [];
    const result = arrayIntersection(nums1, nums2);
    expect(result).toEqual([]);
  });

  test("Test with arrays containing duplicate elements", () => {
    const nums1 = [4, 9, 5, 9];
    const nums2 = [9, 4, 9, 8, 4];
    const result = arrayIntersection(nums1, nums2);
    expect(result.sort()).toEqual([4, 9].sort()); // Sort both arrays
  });
});
