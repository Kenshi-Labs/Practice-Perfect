function removeDuplicates(nums: (number | string)[]): (number | string)[] {
    if (nums.length === 0) return nums;

    let index = 1;
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] !== nums[i - 1]) {
            nums[index] = nums[i];
            index++;
        }
    }

    for (let i = index; i < nums.length; i++) {
        nums[i] = '_';
    }

    return nums;
}

console.log(removeDuplicates([1,1,2]));



export default removeDuplicates;