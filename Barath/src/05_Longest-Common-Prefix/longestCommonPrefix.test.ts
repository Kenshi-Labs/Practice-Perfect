import longestCommonPrefix from './longestCommonPrefix';

test('Test with an array of strings with a common prefix', () => {
    expect(longestCommonPrefix(["flower","flow","flight"])).toBe("fl");
});

test('Test with an array of strings with no common prefix', () => {
    expect(longestCommonPrefix(["dog","racecar","car"])).toBe("");
});

test('Test with an array containing an empty string', () => {
    expect(longestCommonPrefix(["","b"])).toBe("");
});

test('Test with an array containing only one string', () => {
    expect(longestCommonPrefix(["flower"])).toBe("flower");
});
