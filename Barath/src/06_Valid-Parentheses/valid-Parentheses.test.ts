import  isValid  from './valid-Parentheses'

test('valid string of parentheses', () => {
    expect(isValid("()[]{}")).toBe(true);
});

test('invalid string of parentheses', () => {
    expect(isValid("(]")).toBe(false);
});

test('empty string', () => {
    expect(isValid("")).toBe(true);
});

test('string containing only opening brackets', () => {
    expect(isValid("(({{[[")).toBe(false);
});
