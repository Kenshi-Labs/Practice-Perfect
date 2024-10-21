function isPalindrome(s: string): boolean {
  let filteredString = "";
  for (let i = 0; i < s.length; i++) {
    if (/[a-zA-Z0-9]/.test(s[i])) {
      filteredString += s[i].toLowerCase();
    }
  }

  let left = 0;
  let right = filteredString.length - 1;

  while (left < right) {
    if (filteredString[left] !== filteredString[right]) {
      return false;
    }
    left++;
    right--;
  }

  return true;
}
