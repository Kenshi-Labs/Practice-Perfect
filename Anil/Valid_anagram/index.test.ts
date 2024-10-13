import anagram from "./index";

test("Test with valid anagrams", () => {
  expect(anagram("anagram", "nagaram")).toBe(true);
});

test("Test with strings of different lengths", () => {
  expect(anagram("anil", "sunil")).toBe(false);
});

test("Test with strings containing the same letters but different frequencies", () => {
  expect(anagram("aabbcc", "aabbc")).toBe(false);
});

test("Test with empty strings", () => {
  expect(anagram("", "")).toBe(true);
});
