import { compressString } from "./index";

test("Test with a string that can be compressed", () => {
  const input = "aabcccccaaa";
  const output = "a2b1c5a3";
  expect(compressString(input)).toEqual(output);
});

test("Test with a string that cannot be compressed", () => {
  const input = "abcd";
  expect(compressString(input)).toEqual(input);
});

test("Test with an empty string", () => {
  const input = "";
  expect(compressString(input)).toEqual(input);
});

test("Test with a string containing no repeated characters", () => {
  const input = "abc";
  expect(compressString(input)).toEqual(input);
});

test("Test with a string containing all repeated characters", () => {
  const input = "aaaa";
  const output = "a4";
  expect(compressString(input)).toEqual(output);
});

test("Test with a single character string", () => {
  const input = "a";
  expect(compressString(input)).toEqual(input);
});
