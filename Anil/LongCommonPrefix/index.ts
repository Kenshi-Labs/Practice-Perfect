export default function longestCommonPrefix(strs: any) {
  if (strs.length === 0) return ""; // Handle empty array case
  let prefix = strs[0]; // Start with the first string as the prefix

  for (let i = 1; i < strs.length; i++) {
    // Compare prefix with each string in the array and shorten it if necessary
    while (strs[i].indexOf(prefix) !== 0) {
      prefix = prefix.substring(0, prefix.length - 1);
      if (prefix === "") return ""; // No common prefix
    }
  }

  return prefix;
}
