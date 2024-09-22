// Method to reverse, split and join method for reverse String, it gives input to function by passing parameter str

function string_reversed(str: string): string {
  // return str.split("").reverse().join("");

  let reverse_str = "";

  for (let i = str.length - 1; i >= 0; i--) {
    reverse_str += str[i];
  }
  return reverse_str;
}
