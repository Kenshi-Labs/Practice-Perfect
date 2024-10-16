import  longestCommonPrefix  from './longestCommonPrefix';

describe('longestCommonPrefix', () => {
  test('should find the common prefix in an array of strings', () => {
    const strs = ["flower", "flow", "flight"];
    expect(longestCommonPrefix(strs)).toBe("fl");
  });

  test('should return an empty string when there is no common prefix', () => {
    const strs = ["dog", "racecar", "car"];
    expect(longestCommonPrefix(strs)).toBe("");
  });

  test('should return an empty string when an array contains an empty string', () => {
    const strs = ["", "b", "c"];
    expect(longestCommonPrefix(strs)).toBe("");
  });

  test('should return the single string when array contains only one string', () => {
    const strs = ["single"];
    expect(longestCommonPrefix(strs)).toBe("single");
  });
});
