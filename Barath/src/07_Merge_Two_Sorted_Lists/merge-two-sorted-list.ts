function mergeTwoSortedList(list1: number[], list2: number[]): number[] {
  const result: number[] = [];
  let i = 0,
    j = 0;

  while (i < list1.length && j < list2.length) {
    if (list1[i] <= list2[j]) {
      result.push(list1[i++]);
    } else {
      result.push(list2[j++]);
    }
  }

  return result.concat(list1.slice(i)).concat(list2.slice(j));
}

console.log(mergeTwoSortedList([1, 2, 4], [1, 3, 4]));

export default mergeTwoSortedList;
