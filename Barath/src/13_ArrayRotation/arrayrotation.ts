export default function rotateArray(nums: number[], k: number): void {
    const n = nums.length;
    if (n === 0) return;
  
    k = k % n;
  
    for (let i = 0; i < k; i++) {
      nums.unshift(nums.pop()!); 
    }
  }
  
  const nums = [1, 2, 3, 4, 5];
  const k = 2;
  
  rotateArray(nums, k);
  console.log(nums); 
  