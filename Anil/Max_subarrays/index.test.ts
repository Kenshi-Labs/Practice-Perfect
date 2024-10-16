import maxSubArray from "./index";

test("Test with an array containing both positive and negative numbers", () => {
  expect(maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4])).toBe(6);
});

test("Test with an array containing only negative numbers", () => {
  expect(maxSubArray([-3, -4, -1, -2])).toBe(-1);
});

test("Test with an array containing only positive numbers", () => {
  expect(maxSubArray([1, 2, 3, 4])).toBe(10);
});

test("Test with an array containing a single element", () => {
  expect(maxSubArray([5])).toBe(5);
});
