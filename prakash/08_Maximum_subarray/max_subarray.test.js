const maxArray = require("./max_subarray");

describe("Maximum Sub Array", () => {
  test("Test with an array containing both positive and negative numbers", () => {
    const nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4];
    const output = 6;

    expect(maxArray(nums)).toBe(output);
  });

  test("Test with an array containing only negative numbers", () => {
    const nums = [-8, -3, -6, -2, -5, -4];
    const output = -2;
    expect(maxArray(nums)).toBe(output);
  });

  test("Test with an array containing only positive numbers", () => {
    const nums = [1, 2, 3, 4, 5];
    const output = 15;
    expect(maxArray(nums)).toBe(output);
  });

  test("Test with an array containing a single element", () => {
    const nums = [7];
    const output = 7;
    expect(maxArray(nums)).toBe(output);
  });
});
