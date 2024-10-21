function isPalindrome(s) {
  var filteredString = "";
  for (var i = 0; i < s.length; i++) {
    if (/[a-zA-Z0-9]/.test(s[i])) {
      filteredString += s[i].toLowerCase();
    }
  }
  var left = 0;
  var right = filteredString.length - 1;
  while (left < right) {
    if (filteredString[left] !== filteredString[right]) {
      return false;
    }
    left++;
    right--;
  }
  return true;
}
module.exports = isPalindrome;
