const removeDuplicates = require("./RemoveDuplicates");
// removeDuplicates.test.ts

describe("Remove Duplicates from Sorted Array", () => {
  test("Test with an array containing duplicates", () => {
    const nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
    const length = removeDuplicates(nums);
    expect(length).toBe(5); // Unique elements are [0, 1, 2, 3, 4]
    expect(nums.slice(0, length)).toEqual([0, 1, 2, 3, 4]);
  });

  test("Test with an array with no duplicates", () => {
    const nums = [1, 2, 3, 4, 5];
    const length = removeDuplicates(nums);
    expect(length).toBe(5); // No duplicates
    expect(nums.slice(0, length)).toEqual([1, 2, 3, 4, 5]);
  });

  test("Test with an empty array", () => {
    const nums = [];
    const length = removeDuplicates(nums);
    expect(length).toBe(0);
  });

  test("Test with an array containing all identical elements", () => {
    const nums = [1, 1, 1, 1, 1];
    const length = removeDuplicates(nums);
    expect(length).toBe(1); // Only 1 unique element
    expect(nums.slice(0, length)).toEqual([1]);
  });
});
