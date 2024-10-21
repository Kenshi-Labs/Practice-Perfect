import removeDuplicates  from './removeDuplicates'; 

test('array containing duplicates', () => {
    const nums: (number | string)[] = [1, 1, 2];
    const result = removeDuplicates(nums);
    expect(result).toEqual([1, 2, '_']);
});

test('array with no duplicates', () => {
    const nums: (number | string)[] = [1, 2, 3];
    const result = removeDuplicates(nums);
    expect(result).toEqual([1, 2, 3]);
});

test('empty array', () => {
    const nums: (number | string)[] = [];
    const result = removeDuplicates(nums);
    expect(result).toEqual([]);
});

test('array containing all identical elements', () => {
    const nums: (number | string)[] = [1, 1, 1, 1];
    const result = removeDuplicates(nums);
    expect(result).toEqual([1, '_', '_', '_']);
});

