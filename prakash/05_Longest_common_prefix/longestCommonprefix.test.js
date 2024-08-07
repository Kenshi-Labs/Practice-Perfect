const CommonPrefix = require("./longestCommonprefix");

describe("longestCommonPrefix", () => {
  test("Test with an array of strings with a common prefix", () => {
    const input = ["flower", "flow", "flight"];
    const output = "fl";

    expect(CommonPrefix(input)).toBe(output);
  });
  test("Test with an array of strings with no common prefix", () => {
    const input = ["Anil", "Naveen", "Bharath"];
    const output = "";

    expect(CommonPrefix(input)).toBe(output);
  });

  test("Test with an array containing an empty string", () => {
    const input = ["", "abc", "def"];
    const output = "";
    expect(CommonPrefix(input)).toBe(output);
  });

  test("Test with an array containing only one string", () => {
    const input = ["One"];
    const output = "One";
    expect(CommonPrefix(input)).toBe(output);
  });
});
