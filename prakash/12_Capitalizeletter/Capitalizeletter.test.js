const capitalizeFirstLetter = require("./Capitalizeletter");
describe("capitalizeFirstLetter", () => {
  test("Test with a sentence containing multiple words", () => {
    expect(capitalizeFirstLetter("the quick brown fox")).toBe(
      "The Quick Brown Fox"
    );
  });

  test("Test with a single word", () => {
    expect(capitalizeFirstLetter("hello")).toBe("Hello");
  });

  test("Test with a string containing numbers and special characters", () => {
    expect(capitalizeFirstLetter("123 hello! world")).toBe("123 Hello! World");
  });

  test("Test with an empty string", () => {
    expect(capitalizeFirstLetter("")).toBe("");
  });
});
