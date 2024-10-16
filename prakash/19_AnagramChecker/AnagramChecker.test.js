// anagramChecker.test.ts
const areAnagrams = require("./AnagramChecker");
describe("Anagram Checker", () => {
  test("Test with valid anagrams", () => {
    expect(areAnagrams("listen", "silent")).toBe(true);
    expect(areAnagrams("evil", "vile")).toBe(true);
    expect(areAnagrams("elbow", "below")).toBe(true);
  });

  test("Test with strings that are not anagrams", () => {
    expect(areAnagrams("hello", "world")).toBe(false);
    expect(areAnagrams("cat", "dog")).toBe(false);
    expect(areAnagrams("abc", "def")).toBe(false);
  });

  test("Test with strings of different lengths", () => {
    expect(areAnagrams("abc", "abcd")).toBe(false);
    expect(areAnagrams("a", "aa")).toBe(false);
  });

  test("Test with empty strings", () => {
    expect(areAnagrams("", "")).toBe(true);
  });
});
