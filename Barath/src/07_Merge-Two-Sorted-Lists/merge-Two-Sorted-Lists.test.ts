import { mergeSortedLists } from './merge-Two-Sorted-Lists';

describe('mergeSortedLists', () => {
    test('merge two non-empty sorted lists', () => {
        const arr1 = [1, 2, 4];
        const arr2 = [1, 3, 4];
        expect(mergeSortedLists(arr1, arr2)).toEqual([1, 1, 2, 3, 4, 4]);
    });

    test('merge one empty list and one non-empty list', () => {
        const arr1: number[] = [];
        const arr2 = [1, 3, 4];
        expect(mergeSortedLists(arr1, arr2)).toEqual([1, 3, 4]);
    });

    test('merge two empty lists', () => {
        const arr1: number[] = [];
        const arr2: number[] = [];
        expect(mergeSortedLists(arr1, arr2)).toEqual([]);
    });

    test('merge lists of different lengths', () => {
        const arr1 = [1, 2, 4, 5, 6];
        const arr2 = [1, 3, 4];
        expect(mergeSortedLists(arr1, arr2)).toEqual([1, 1, 2, 3, 4, 4, 5, 6]);
    });

    test('merge one non-empty list and one empty list', () => {
        const arr1 = [1, 2, 3];
        const arr2: number[] = [];
        expect(mergeSortedLists(arr1, arr2)).toEqual([1, 2, 3]);
    });

    test('merge two lists where all elements of one are greater than the other', () => {
        const arr1 = [5, 6, 7];
        const arr2 = [1, 2, 3];
        expect(mergeSortedLists(arr1, arr2)).toEqual([1, 2, 3, 5, 6, 7]);
    });
});
