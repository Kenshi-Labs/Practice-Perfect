import mergeTwoSortedList from "./merge-two-sorted-list";

describe('mergeTwoSortedList', () => {
    it('should merge two non-empty sorted lists', () => {
        expect(mergeTwoSortedList([1, 2, 4], [1, 3, 4])).toEqual([1, 1, 2, 3, 4, 4]);
    });

    it('should merge one empty list and one non-empty list', () => {
        expect(mergeTwoSortedList([], [1, 3, 4])).toEqual([1, 3, 4]);
    });

    it('should merge two empty lists', () => {
        expect(mergeTwoSortedList([], [])).toEqual([]);
    });

    it('should merge lists of different lengths', () => {
        expect(mergeTwoSortedList([1, 2, 4], [1, 3, 4, 5])).toEqual([1, 1, 2, 3, 4, 4, 5]);
    });
});

