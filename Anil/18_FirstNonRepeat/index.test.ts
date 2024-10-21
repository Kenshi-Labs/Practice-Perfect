import { firstNonRepeatingChar } from "./index";

test("Test with a string containing a non-repeating character", () => {
  const input = "aabbcdeeff";
  expect(firstNonRepeatingChar(input)).toEqual("c");
});

test("Test with a string where all characters repeat", () => {
  const input = "aabbcc";
  expect(firstNonRepeatingChar(input)).toBeNull(); // No non-repeating characters
});

test("Test with an empty string", () => {
  const input = "";
  expect(firstNonRepeatingChar(input)).toBeNull(); // No characters to check
});

test("Test with a string containing only one character", () => {
  const input = "x";
  expect(firstNonRepeatingChar(input)).toEqual("x"); // Single character is non-repeating by default
});

test("Test with a string containing multiple non-repeating characters", () => {
  const input = "abcabcde";
  expect(firstNonRepeatingChar(input)).toEqual("d"); // 'd' is the first non-repeating character
});
