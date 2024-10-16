import climbStairs from "./index";

test("Test with n = 1", () => {
  expect(climbStairs(1)).toBe(1);
});

test("Test with n = 2", () => {
  expect(climbStairs(2)).toBe(2);
});

test("Test with n = 5", () => {
  expect(climbStairs(5)).toBe(8);
});

test("Test with n = 0", () => {
  expect(climbStairs(0)).toBe(1);
});
