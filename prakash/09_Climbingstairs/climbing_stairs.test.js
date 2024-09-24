const climbingStairs = require("./climbing_stairs");

describe(" climbing Stairs", () => {
  test("Test with n = 1", () => {
    const input = 1;
    const output = 1;

    expect(climbingStairs(input)).toBe(output);
  });

  test("Test with n = 2", () => {
    const input = 2;
    const output = 2;

    expect(climbingStairs(input)).toBe(output);
  });

  test("Test with n = 5", () => {
    const input = 5;
    const output = 8;

    expect(climbingStairs(input)).toBe(output);
  });

  test("Test with n = 0", () => {
    const input = 0;
    const output = 0;

    expect(climbingStairs(input)).toBe(output);
  });
});
