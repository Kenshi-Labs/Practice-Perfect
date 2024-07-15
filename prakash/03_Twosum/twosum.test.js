const twosum = require("./twosum");

describe("two_sum", () => {
  // 1 Test with a valid input where a solution exists

  test("should return indices of the two numbers that add up to the target", () => {
    expect(twosum([2, 7, 11, 15], 9)).toStrictEqual([0, 1]);
  });

  // 2 Test with no valid solution

  test("should throw an error if no valid solution exists", () => {
    expect(() => twosum([2, 7, 11, 15], 10)).toThrow("No two sum solution");
  });

  // 3 Test with multiple valid solutions (return any one)

  test("should return indices of one valid solution when multiple solutions exist", () => {
    const result = twosum([3, 2, 4, 3], 6);
    expect([
      [0, 3],
      [1, 2],
    ]).toContainEqual(result);
  });

  // 4 Test with an array of length 2

  test("should return indices for an array of length 2", () => {
    expect(twosum([1, 2], 3)).toEqual([0, 1]);
  });
});
