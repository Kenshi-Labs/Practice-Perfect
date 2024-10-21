// capitalizeFirstLetter.test.ts

import { capitalizeFirstLetter } from './capitalizeFirstLetter';

describe('capitalizeFirstLetter', () => {

    // Test with a sentence containing multiple words
    test('should capitalize the first letter of each word in a sentence', () => {
        const input = "the quick brown fox";
        const result = capitalizeFirstLetter(input);
        expect(result).toBe("The Quick Brown Fox");
    });

    // Test with a single word
    test('should capitalize the first letter of a single word', () => {
        const input = "hello";
        const result = capitalizeFirstLetter(input);
        expect(result).toBe("Hello");
    });

    // Test with an empty string
    test('should return an empty string when input is empty', () => {
        const input = "";
        const result = capitalizeFirstLetter(input);
        expect(result).toBe("");
    });
});
