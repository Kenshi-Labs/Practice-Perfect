const longestCommonPrefix = require("./LongestCommonPrefix");

describe("Longest Common Prefix", () => {
  test("Test with an array of strings with a common prefix", () => {
    const input = ["flower", "flow", "flight"];
    const output = "fl";
    expect(longestCommonPrefix(input)).toBe(output);
  });

  test("Test with an array of strings with no common prefix", () => {
    const input = ["dog", "racecar", "car"];
    const output = "";
    expect(longestCommonPrefix(input)).toBe(output);
  });

  test("Test with an array containing an empty string", () => {
    const input = ["", "b"];
    const output = "";
    expect(longestCommonPrefix(input)).toBe(output);
  });

  test("Test with an array containing only one string", () => {
    const input = ["single"];
    const output = "single";
    expect(longestCommonPrefix(input)).toBe(output);
  });
});
