import reverseString from "./index";

test("Test with a non-empty string", () => {
  expect(reverseString("hello")).toBe("olleh");
});

test("Test with an empty string", () => {
  expect(reverseString("")).toBe("");
});

test("Test with a string containing spaces", () => {
  expect(reverseString("hello world")).toBe("dlrow olleh");
});

test("Test with a palindrome", () => {
  expect(reverseString("dad")).toBe("dad");
});
