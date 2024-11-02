function reverseString(input) {
    let reverse_str = "";

  for (let i = input.length - 1; i >= 0; i--) {
    reverse_str += input[i];
  }
  return reverse_str;
}

module.exports = reverseString;
