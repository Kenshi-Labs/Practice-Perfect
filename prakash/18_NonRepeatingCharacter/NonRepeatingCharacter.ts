function firstNonRepeatingChar(s: string): string | null {
  const charCount: { [key: string]: number } = {}; // Step 1: Create a map to store character frequencies

  // Step 2: Count the frequency of each character
  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    charCount[char] = (charCount[char] || 0) + 1;
  }

  // Step 3: Find the first character that appears only once
  for (let i = 0; i < s.length; i++) {
    if (charCount[s[i]] === 1) {
      return s[i];
    }
  }

  return null;
}
