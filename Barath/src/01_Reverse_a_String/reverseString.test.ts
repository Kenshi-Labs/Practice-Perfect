import  reverseString from './reverseString'; 

test('reverse a non-empty string', () => {
    expect(reverseString("hello")).toBe("olleh");
});

test('reverse an empty string', () => {
    expect(reverseString("")).toBe("");
});

test('reverse a string containing spaces', () => {
    expect(reverseString("hello world")).toBe("dlrow olleh");
});

test('reverse a palindrome', () => {
    expect(reverseString("madam")).toBe("madam");
});
