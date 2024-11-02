import longestCommonPrefix from "./index";

test("array of strings with a common prefix", () => {
  const strs = ["flower", "flow", "flight"];
  const result = longestCommonPrefix(strs);
  expect(result).toBe("fl");
});

test("array of strings with no common prefix", () => {
  const strs = ["dog", "racecar", "car"];
  const result = longestCommonPrefix(strs);
  expect(result).toBe("");
});

test("array containing an empty string", () => {
  const strs = ["", "b", "c"];
  const result = longestCommonPrefix(strs);
  expect(result).toBe("");
});

test("array containing only one string", () => {
  const strs = ["single"];
  const result = longestCommonPrefix(strs);
  expect(result).toBe("single");
});
