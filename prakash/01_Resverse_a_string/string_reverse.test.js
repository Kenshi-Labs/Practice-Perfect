// import file from string_reverse.js
const string_reverse = require("./string_reverse.js");

// function where test case performed
test("Reverse String is", () => {
  //create a array to perform multiple test cases
  const test = [
    //Test with a non-empty string
    { input: "hello", expected: "olleh" },

    //Test with empty string
    { input: "", expected: "" },

    // Test case with Extra Spaces
    { input: "h e l l o", expected: "o l l e h" },

    // Test case for palindrome
    { input: "madam", expected: "madam" },
  ];

  // Method array (forEach) is used to performed the each test cases to pass the result and time
  test.forEach((test, index) => {
    //Give input to function string_reverse and stored in variable result
    const result = string_reverse(test.input);

    //Compare the result with expected result and pass the result
    expect(result).toBe(test.expected);

    // it provides serial no. for each test case
    console.log(`Test case ${index + 1} pass`);
  });
});
