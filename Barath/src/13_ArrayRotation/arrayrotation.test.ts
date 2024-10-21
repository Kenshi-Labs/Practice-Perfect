import rotateArray from './arrayrotation';

describe('rotateArray', () => {
    test('should rotate a non-empty array by a positive number of steps', () => {
        const nums = [1, 2, 3, 4, 5];
        rotateArray(nums, 2);
        expect(nums).toEqual([4, 5, 1, 2, 3]);
    });

    test('should rotate the array by its length (no change)', () => {
        const nums = [1, 2, 3, 4, 5];
        rotateArray(nums, 5);
        expect(nums).toEqual([1, 2, 3, 4, 5]);
    });

    test('should rotate the array when rotation is greater than array length', () => {
        const nums = [1, 2, 3, 4, 5];
        rotateArray(nums, 7); // 7 % 5 = 2
        expect(nums).toEqual([4, 5, 1, 2, 3]);
    });

    test('should handle an empty array', () => {
        const nums: number[] = [];
        rotateArray(nums, 3);
        expect(nums).toEqual([]);
    });
});
