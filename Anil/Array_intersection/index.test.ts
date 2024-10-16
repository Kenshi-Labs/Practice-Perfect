import commonElements from "./index";

test("Test with arrays having common elements", () => {
  const nums1 = [1, 2, 2, 1];
  const nums2 = [2, 2, 3];
  expect(commonElements(nums1, nums2)).toEqual([2]);
});

test("Test with arrays having no common elements", () => {
  const nums1 = [1, 4, 5];
  const nums2 = [2, 6, 3];
  expect(commonElements(nums1, nums2)).toEqual([]);
});

test("Test with empty arrays", () => {
  const nums1: any = [];
  const nums2: any = [];
  expect(commonElements(nums1, nums2)).toEqual([]);
});

test("Test with arrays containing duplicate elements", () => {
  const nums1 = [1, 2, 2, 3, 4];
  const nums2 = [2, 2, 4, 4];
  expect(commonElements(nums1, nums2)).toEqual([2, 4]);
});
