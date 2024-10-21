export function arrayIntersection(nums1: number[], nums2: number[]): number[] {
    const set1 = new Set(nums1);
    const set2 = new Set(nums2);
    const result: number[] = [];

    set1.forEach((value) => {
        if (set2.has(value)) {
            result.push(value);
        }
    });

    return result;
}

const nums1 = [1, 2, 2, 1];
const nums2 = [2, 2, 3];
console.log(arrayIntersection(nums1, nums2));