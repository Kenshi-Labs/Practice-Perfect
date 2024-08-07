function IsValid(s: string): boolean {
  function removePairs(str: string) {
    if (str.length === 0) return str;

    const regex = /(\(\)|\[\]|\{\})/g;

    let Newstr = str.replace(regex, "");

    if (Newstr === str) return str;

    return removePairs(Newstr);
  }
  return removePairs(s) === "";
}
