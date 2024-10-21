export function isPalindrome(input: string): boolean {
  const cleaned = input.replace(/[^a-z0-9]/gi, "").toLowerCase();
  return cleaned === cleaned.split("").reverse().join("");
}
