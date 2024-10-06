const isPalindrome = require("./PalindromeCheck");

describe("isPalindrome", () => {
  // Test 1: Valid palindrome containing spaces and punctuation
  test("Test with a valid palindrome containing spaces and punctuation", () => {
    const input = "A man, a plan, a canal: Panama";
    expect(isPalindrome(input)).toBe(true);
  });

  // Test 2: Non-palindrome
  test("Test with a non-palindrome", () => {
    const input = "hello";
    expect(isPalindrome(input)).toBe(false);
  });

  // Test 3: Empty string
  test("Test with an empty string", () => {
    const input = "";
    expect(isPalindrome(input)).toBe(true); // An empty string is considered a palindrome
  });

  // Test 4: String with only non-alphanumeric characters
  test("Test with a string containing only non-alphanumeric characters", () => {
    const input = "!!@@##";
    expect(isPalindrome(input)).toBe(true); // An empty string after removing non-alphanumeric characters
  });
});
