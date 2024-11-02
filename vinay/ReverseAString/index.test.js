// import { reverseString } from "./index";

// /**
//  * Unit Tests for reverse Strings
//  **/

// function runTests() {
//     // Non-empty string
//     console.log(reverseString("hello") === "olleh");
//     // Empty string
//     console.log(reverseString("") === "");
//     // String with spaces
//     console.log(reverseString("hello world") === "dlrow olleh");
//     // Palindrome
//     console.log(reverseString("madam") === "madam");
//     console.log("Now All tests passed!");
// }

// // Tests
// runTests();


const reverseString = require('./index');

describe('reverseString', () => {
  
  it('Test with a non-empty string', () => {
    expect(reverseString('hello')).toBe('olleh');
  });

  it('Test with an empty string', () => {
    expect(reverseString('')).toBe('');
  });

  it('Test with a string containing spaces', () => {
    expect(reverseString('hello world')).toBe('dlrow olleh');
  });

  it('Test with a palindrome', () => {
    expect(reverseString('madam')).toBe('madam');
  });

});

