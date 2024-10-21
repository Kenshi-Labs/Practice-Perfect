import rotate from "./index"; // Assuming the function is in 'index.js'

test("Test with a non-empty array and positive rotation", () => {
  const nums = [1, 2, 3, 4, 5];
  rotate(nums, 2);
  expect(nums).toEqual([4, 5, 1, 2, 3]);
});

test("Test with rotation equal to array length", () => {
  const nums = [1, 2, 3, 4, 5];
  rotate(nums, 5); // Rotating by the array length should result in the same array
  expect(nums).toEqual([1, 2, 3, 4, 5]);
});

test("Test with rotation greater than array length", () => {
  const nums = [1, 2, 3, 4, 5];
  rotate(nums, 7); // Equivalent to rotating by 7 % 5 = 2
  expect(nums).toEqual([4, 5, 1, 2, 3]);
});

test("Test with an empty array", () => {
  const nums: any = [];
  rotate(nums, 3); // No change expected
  expect(nums).toEqual([]);
});
