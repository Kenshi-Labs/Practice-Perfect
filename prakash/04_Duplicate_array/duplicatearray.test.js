const duplicatearray = require("./duplicatearray");

describe("RemovesDuplicates", () => {

  // 1. Test with an array containing duplicates

  test("Test with an array containing duplicates", () => {
    const nums = [1, 1, 2];
    const length = duplicatearray(nums);
    expect(length).toBe(2);
    expect(nums.slice(0, length)).toEqual([1, 2]);
  });

  // 2. Test with an array with no duplicates

  test("Test with an array with no duplicates", () => {
    const nums = [1, 2, 3];
    const length = duplicatearray(nums);
    expect(length).toBe(3);
    expect(nums.slice(0, length)).toEqual([1, 2, 3]);
  })

  // 3. Test with an empty array

  test('Test with an empty array', () => {
    const nums = [];
    const length = duplicatearray(nums);
    expect(length).toBe(0);
    expect(nums).toEqual([]);
  })

  // 4. Test with an array containing all identical elements

  test('Test with an array containing all identical elements', () => {
    const nums = [2, 2, 2, 2, 2];
    const length = duplicatearray(nums);
    expect(length).toBe(1);
    expect(nums.splice(0, length)).toEqual([2]);
  })
});
