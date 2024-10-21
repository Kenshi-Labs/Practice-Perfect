export function firstNonRepeatingChar(input: string): string | null {
  const charCount: { [key: string]: number } = {};

  // Count the occurrences of each character
  for (const char of input) {
    charCount[char] = (charCount[char] || 0) + 1;
  }

  // Find the first character that appears only once
  for (const char of input) {
    if (charCount[char] === 1) {
      return char;
    }
  }

  // If no non-repeating character found, return null
  return null;
}
