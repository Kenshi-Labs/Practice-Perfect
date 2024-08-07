function IsValid(s) {
  function removePairs(str) {
    if (str.length === 0) return str;
    var regex = /(\(\)|\[\]|\{\})/g;
    var Newstr = str.replace(regex, "");
    if (Newstr === str) return str;
    return removePairs(Newstr);
  }
  return removePairs(s) === "";
}

module.exports = IsValid;
