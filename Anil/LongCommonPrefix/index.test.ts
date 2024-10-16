import longestCommonPrefix from "./index";

test("Test with an array of strings with a common prefix", () => {
  const strs = ["flower", "flow", "flight"];
  expect(longestCommonPrefix(strs)).toBe("fl");
});

test("Test with an array of strings with no common prefix", () => {
  const strs = ["dog", "racecar", "car"];
  expect(longestCommonPrefix(strs)).toBe("");
});

test("Test with an array containing an empty string", () => {
  const strs = ["flower", "flow", ""];
  expect(longestCommonPrefix(strs)).toBe("");
});

test("Test with an array containing only one string", () => {
  const strs = ["single"];
  expect(longestCommonPrefix(strs)).toBe("single");
});
