// arrayIntersection.test.ts
import { arrayIntersection } from './arrayintersection';

describe('arrayIntersection', () => {
    test('should return common elements between two arrays', () => {
        const nums1 = [1, 2, 2, 1];
        const nums2 = [2, 2, 3];
        expect(arrayIntersection(nums1, nums2)).toEqual([2]);
    });

    test('should return an empty array when there are no common elements', () => {
        const nums1 = [1, 2, 3];
        const nums2 = [4, 5, 6];
        expect(arrayIntersection(nums1, nums2)).toEqual([]);
    });

    test('should return an empty array when both arrays are empty', () => {
        const nums1: number[] = [];
        const nums2: number[] = [];
        expect(arrayIntersection(nums1, nums2)).toEqual([]);
    });

    test('should handle arrays with duplicates and return unique common elements', () => {
        const nums1 = [4, 9, 5, 4];
        const nums2 = [9, 4, 9, 8, 4];
        expect(arrayIntersection(nums1, nums2)).toEqual([4, 9]);
    });

    test('should handle arrays with one element', () => {
        const nums1 = [1];
        const nums2 = [1];
        expect(arrayIntersection(nums1, nums2)).toEqual([1]);
    });

    test('should return an empty array when one of the arrays is empty', () => {
        const nums1 = [1, 2, 3];
        const nums2: number[] = [];
        expect(arrayIntersection(nums1, nums2)).toEqual([]);
    });
});
