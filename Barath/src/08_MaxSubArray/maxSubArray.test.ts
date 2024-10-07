import maxSubArray from './maxSubArray';

test('Test with mixed positive and negative numbers', () => {
    const nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4];
    expect(maxSubArray(nums)).toBe(6);
});

test('Test with only negative numbers', () => {
    const nums = [-3, -5, -7, -1];
    expect(maxSubArray(nums)).toBe(-1);
});

test('Test with only positive numbers', () => {
    const nums = [2, 3, 5, 7];
    expect(maxSubArray(nums)).toBe(17);
});

test('Test with single element array', () => {
    const nums = [5];
    expect(maxSubArray(nums)).toBe(5);
});
