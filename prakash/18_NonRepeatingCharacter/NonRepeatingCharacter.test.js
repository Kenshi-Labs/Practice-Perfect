const firstNonRepeatingChar = require("./NonRepeatingCharacter");
describe("firstNonRepeatingChar", () => {
  // Test 1: String containing a non-repeating character
  test("Test with a string containing a non-repeating character", () => {
    const input = "aabbcdeeff";
    expect(firstNonRepeatingChar(input)).toBe("c");
  });

  // Test 2: String where all characters repeat
  test("Test with a string where all characters repeat", () => {
    const input = "aabbcc";
    expect(firstNonRepeatingChar(input)).toBeNull();
  });

  // Test 3: Empty string
  test("Test with an empty string", () => {
    const input = "";
    expect(firstNonRepeatingChar(input)).toBeNull();
  });

  // Test 4: String containing only one character
  test("Test with a string containing only one character", () => {
    const input = "x";
    expect(firstNonRepeatingChar(input)).toBe("x");
  });
});
