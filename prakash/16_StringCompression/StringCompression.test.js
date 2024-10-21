const compressString = require("./StringCompression");

describe("compressString", () => {
  // Test 1: String that can be compressed
  test("Test with a string that can be compressed", () => {
    const input = "aabcccccaaa";
    const output = "a2b1c5a3";
    expect(compressString(input)).toBe(output);
  });

  // Test 2: String that cannot be compressed (should return original)
  test("Test with a string that cannot be compressed (should return original)", () => {
    const input = "abcd";
    const output = "abcd";
    expect(compressString(input)).toBe(output);
  });

  // Test 3: Empty string
  test("Test with an empty string", () => {
    const input = "";
    const output = "";
    expect(compressString(input)).toBe(output);
  });

  // Test 4: String with no repeated characters
  test("Test with a string containing no repeated characters", () => {
    const input = "abcde";
    const output = "abcde";
    expect(compressString(input)).toBe(output);
  });
});
