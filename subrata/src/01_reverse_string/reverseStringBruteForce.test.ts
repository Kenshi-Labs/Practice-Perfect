import reverseStringBruteForce from "./reverseStringBruteForce";

describe("reverseStringBruteForce", () => {
  // 1. Test with a non-empty string
  it("should reverse a non-empty string", () => {
    const input = "hello";
    const output = reverseStringBruteForce(input);
    expect(output).toBe("olleh");
  });

  // 2. Test with an empty string
  it("should return an empty string when the input is empty", () => {
    const input = "";
    const output = reverseStringBruteForce(input);
    expect(output).toBe("");
  });

  // 3. Test with a string containing spaces
  it("should reverse a string containing spaces", () => {
    const input = "hello world";
    const output = reverseStringBruteForce(input);
    expect(output).toBe("dlrow olleh");
  });

  // 4. Test with a palindrome
  it("should reverse a palindrome", () => {
    const input = "madam";
    const output = reverseStringBruteForce(input);
    expect(output).toBe("madam");
  });
});
