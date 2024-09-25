function ArrayIntersection(
  nums1: number[],
  nums2: number[],
  nums3: number[]
): number[] {
  const set1 = new Set(nums1);
  const resultSet = new Set<number>();

  for (const num of nums2) {
    if (set1.has(num)) {
      resultSet.add(num);
    }
  }
  return Array.from(resultSet);
}
