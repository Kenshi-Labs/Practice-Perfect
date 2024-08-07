const twosum = require("./twosum");

describe("two_sum", () => {
  // 1 Test with a valid input where a solution exists

  test("should return indices of the two numbers that add up to the target", () => {
    expect(twosum([2, 7, 11, 15], 9)).toStrictEqual([0, 1]);
  });

  // 2 Test with no valid solution

  test("return an empty array if no valid solution exists", () => {
    expect(twosum([2, 7, 11, 15], 10)).toEqual([]);
  });

  // 3 Test with multiple valid solutions (return any one)

  test("return any one pair of indices if multiple valid solutions exist", () => {
    const result = twosum([2, 7, 11, 15], 18);
    const validpairs = [
      [1, 2],
      [2, 1],
    ];
    expect(validpairs).toContainEqual(result);
  });

  // 4 Test with an array of length 2

  test("should return indices for an array of length 2", () => {
    expect(twosum([2, 7], 9)).toEqual([0, 1]);
  });
});
