export default function commonElements(nums1: any, nums2: any) {
  // Create sets to remove duplicates
  const set1 = new Set(nums1);
  const set2 = new Set(nums2);

  // Find common elements between the two sets
  const result = [...set1].filter((item) => set2.has(item));

  return result;
}
