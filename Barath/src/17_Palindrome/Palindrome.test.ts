

import { isPalindrome } from './Palindrome';

describe('isPalindrome', () => {
    test('should return true for a valid palindrome with spaces and punctuation', () => {
        const input = "A man, a plan, a canal: Panama";
        expect(isPalindrome(input)).toBe(true);
    });

    test('should return false for a non-palindrome', () => {
        const input = "This is not a palindrome";
        expect(isPalindrome(input)).toBe(false);
    });

    test('should return true for an empty string', () => {
        const input = "";
        expect(isPalindrome(input)).toBe(true);
    });

    test('should return true for a string containing only non-alphanumeric characters', () => {
        const input = "!!!@@@###";
        expect(isPalindrome(input)).toBe(true);
    });

    test('should return true for a palindrome with mixed case characters', () => {
        const input = "No lemon, no melon";
        expect(isPalindrome(input)).toBe(true);
    });
});
