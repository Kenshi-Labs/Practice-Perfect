function areAnagrams(s1, s2) {
  // Step 1: Check if the lengths are different
  if (s1.length !== s2.length) {
    return false;
  }
  // Step 2: Create a map (object) to count letters in s1
  var letterCount = {};
  // Step 3: Count letters in s1
  for (var i = 0; i < s1.length; i++) {
    var char = s1[i];
    letterCount[char] = (letterCount[char] || 0) + 1;
  }
  // Step 4: Subtract the count using letters from s2
  for (var i = 0; i < s2.length; i++) {
    var char = s2[i];
    if (!letterCount[char]) {
      return false; // If a letter in s2 doesn't exist or count doesn't match
    }
    letterCount[char]--;
  }
  // Step 5: Ensure all counts are zero (every letter matched exactly)
  for (var key in letterCount) {
    if (letterCount[key] !== 0) {
      return false;
    }
  }
  return true;
}

module.exports = areAnagrams;
