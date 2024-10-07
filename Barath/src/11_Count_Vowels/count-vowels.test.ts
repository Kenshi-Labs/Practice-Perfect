// countVowels.test.ts

import { countVowels } from './count-vowels';

describe('countVowels', () => {
    
    // Test with a string containing vowels
    test('should count vowels in "Hello, World!"', () => {
        const input = "Hello, World!";
        const result = countVowels(input);
        expect(result).toBe(3);
    });

    // Test with a string containing no vowels
    test('should return 0 when there are no vowels', () => {
        const input = "bcdfg";
        const result = countVowels(input);
        expect(result).toBe(0);
    });

    // Test with an empty string
    test('should return 0 for an empty string', () => {
        const input = "";
        const result = countVowels(input);
        expect(result).toBe(0);
    });

    // Test with a string containing only vowels
    test('should count vowels when only vowels are present', () => {
        const input = "aeiouAEIOU";
        const result = countVowels(input);
        expect(result).toBe(10);
    });
});
