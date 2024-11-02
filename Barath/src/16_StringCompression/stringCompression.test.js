import { stringCompression } from './stringCompression';

describe('stringCompression', () => {
    test('should return compressed string when compression reduces length', () => {
        const input = "aabcccccaaa";
        expect(stringCompression(input)).toBe("a2b1c5a3");
    });

    test('should return original string when compression does not reduce length', () => {
        const input = "abcdef";
        expect(stringCompression(input)).toBe("abcdef");
    });

    test('should return an empty string when input is empty', () => {
        const input = "";
        expect(stringCompression(input)).toBe("");
    });

    test('should return original string when there are no repeated characters', () => {
        const input = "abcd";
        expect(stringCompression(input)).toBe("abcd");
    });

    test('should handle single character strings', () => {
        const input = "a";
        expect(stringCompression(input)).toBe("a");
    });

    test('should handle strings with multiple different repeating patterns', () => {
        const input = "aaabbccccdde";
        expect(stringCompression(input)).toBe("a3b2c4d2e1");
    });
});
