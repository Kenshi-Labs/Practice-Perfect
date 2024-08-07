const {
  MergetwoList,
  createLinkedList,
  linkedListToArray,
} = require("./merge_sorted_list");

describe("MergetwoLists", () => {
  test("mergeTwoLists with two non-empty sorted lists", () => {
    const l1 = createLinkedList([1, 2, 4]);
    const l2 = createLinkedList([1, 3, 4]);
    const mergedList = MergetwoList(l1, l2);
    expect(linkedListToArray(mergedList)).toEqual([1, 1, 2, 3, 4, 4]);
  });

  test("Test with one empty list and one non-empty list", () => {
    const l1 = createLinkedList([]);
    const l2 = createLinkedList([1, 3, 4]);
    const mergedList = MergetwoList(l1, l2);
    expect(linkedListToArray(mergedList)).toEqual([1, 3, 4]);
  });

  test("Test with two empty lists", () => {
    const l1 = createLinkedList([]);
    const l2 = createLinkedList([]);
    const mergedList = MergetwoList(l1, l2);
    expect(linkedListToArray(mergedList)).toEqual([]);
  });

  test("Test with lists of different lengths", () => {
    const l1 = createLinkedList([1, 2, 3]);
    const l2 = createLinkedList([4, 5, 6, 7, 8]);
    const mergedList = MergetwoList(l1, l2);
    expect(linkedListToArray(mergedList)).toEqual([1, 2, 3, 4, 5, 6, 7, 8]);
  });
});
