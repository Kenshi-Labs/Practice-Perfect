function areAnagrams(s1: string, s2: string): boolean {
  if (s1.length !== s2.length) {
    return false;
  }

  const letterCount: { [key: string]: number } = {};

  for (let i = 0; i < s1.length; i++) {
    const char = s1[i];
    letterCount[char] = (letterCount[char] || 0) + 1;
  }

  for (let i = 0; i < s2.length; i++) {
    const char = s2[i];
    if (!letterCount[char]) {
      return false;
    }
    letterCount[char]--;
  }

  for (let key in letterCount) {
    if (letterCount[key] !== 0) {
      return false;
    }
  }

  return true;
}
