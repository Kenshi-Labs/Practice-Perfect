function reverseStringOptimal(input: string): string {
  let arr = input.split("");
  let left = 0;
  let right = arr.length - 1;

  while (left < right) {
    // Swap the characters at the left and right pointers
    [arr[left], arr[right]] = [arr[right], arr[left]];
    left++;
    right--;
  }
  return arr.join("");
}
export default reverseStringOptimal;
