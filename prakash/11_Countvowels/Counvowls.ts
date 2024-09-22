function Countvowels(vowelss: string): number {
  let vowels = new Set(["a", "e", "i", "o", "u"]);
  let count = 0;

  for (const char of vowelss.toLowerCase()) {
    if (vowels.has(char)) {
      count++;
    }
  }
  return count;
}
