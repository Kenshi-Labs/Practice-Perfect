import { isPalindrome } from "./index";

test("Test with a valid palindrome containing spaces and punctuation", () => {
  const input = "A man, a plan, a canal: Panama";
  expect(isPalindrome(input)).toBe(true);
});

test("Test with a non-palindrome", () => {
  const input = "hello";
  expect(isPalindrome(input)).toBe(false);
});

test("Test with an empty string", () => {
  const input = "";
  expect(isPalindrome(input)).toBe(true);
});

test("Test with a string containing only non-alphanumeric characters", () => {
  const input = "!@#$$#@!";
  expect(isPalindrome(input)).toBe(true);
});

test("Test with a single character string", () => {
  const input = "a";
  expect(isPalindrome(input)).toBe(true);
});
