function LongestCommonPrefix(strs) {
  if (strs.length === 0) return "";
  var prefix = strs[0]; // starts with first string;
  for (var i = 1; i < strs.length; i++) {
    while (strs[i].indexOf(prefix) !== 0) {
      prefix = prefix.substring(0, prefix.length - 1);
      if (prefix === "") return "";
    }
  }
  return prefix;
}

module.exports = LongestCommonPrefix;
