import  isAnagram  from './isAnagram'; 

test('valid anagrams', () => {
    expect(isAnagram("anagram", "nagaram")).toBe(true);
});

test('strings of different lengths', () => {
    expect(isAnagram("rat", "car")).toBe(false);
});

test('strings with same letters but different frequencies', () => {
    expect(isAnagram("aabbcc", "aabbc")).toBe(false);
});

test('empty strings', () => {
    expect(isAnagram("", "")).toBe(true);
});
