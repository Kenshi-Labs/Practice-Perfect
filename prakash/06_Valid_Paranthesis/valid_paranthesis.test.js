const isValid = require("./valid_paranthesis");

describe("Valid Parenthesis", () => {
  test("Test with a valid string of parentheses", () => {
    const input = "()[]{}";
    const output = true;

    expect(isValid(input)).toBe(output);
  });

  test("Test with an invalid string of parentheses", () => {
    const input = "(]";
    const output = false;

    expect(isValid(input)).toBe(output);
  });

  test("Test with an empty string", () => {
    const input = "";
    const output = true;

    expect(isValid(input)).toBe(output);
  });

  test("Test with a string containing only opening brackets", () => {
    const input = "([{";
    const output = false;

    expect(isValid(input)).toBe(output);
  });
});
