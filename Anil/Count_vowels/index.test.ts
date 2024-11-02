import countVowels from "./index";

test("Test with a string containing vowels", () => {
  expect(countVowels("Hello, World!")).toBe(3);
});

test("Test with a string containing no vowels", () => {
  expect(countVowels("rhythm")).toBe(0);
});

test("Test with an empty string", () => {
  expect(countVowels("")).toBe(0);
});

test("Test with a string containing only vowels", () => {
  expect(countVowels("aeiou")).toBe(5);
});

test("Test with mixed case vowels", () => {
  expect(countVowels("AeIoU")).toBe(5);
});
