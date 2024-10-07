import  twoSum  from './twoSum';

test('valid input where a solution exists', () => {
    expect(twoSum([2, 7, 11, 15], 9)).toEqual([0, 1]);
});

test('no valid solution', () => {
    expect(() => twoSum([2, 7, 11, 15], 20)).toThrow('No two sum solution');
});

test('multiple valid solutions (return any one)', () => {
    const result = twoSum([3, 2, 4], 6);
    expect(result).toEqual(expect.arrayContaining([1, 2]));
});

test('array of length 2', () => {
    expect(twoSum([1, 2], 3)).toEqual([0, 1]);
});
