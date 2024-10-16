function firstNonRepeatingChar(s) {
  var charCount = {}; // Step 1: Create a map to store character frequencies
  // Step 2: Count the frequency of each character
  for (var i = 0; i < s.length; i++) {
    var char = s[i];
    charCount[char] = (charCount[char] || 0) + 1;
  }
  // Step 3: Find the first character that appears only once
  for (var i = 0; i < s.length; i++) {
    if (charCount[s[i]] === 1) {
      return s[i];
    }
  }
  return null;
}
module.exports = firstNonRepeatingChar;
