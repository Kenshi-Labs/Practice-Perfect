import isValid from "./index";

test("Test with a valid string of parentheses", () => {
  expect(isValid("()[]{}")).toBe(true);
});

test("Test with an invalid string of parentheses", () => {
  expect(isValid("(]")).toBe(false);
});

test("Test with an empty string", () => {
  expect(isValid("")).toBe(true);
});

test("Test with a string containing only opening brackets", () => {
  expect(isValid("(((")).toBe(false);
});
