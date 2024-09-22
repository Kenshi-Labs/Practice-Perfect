// Method to reverse, split and join method for reverse String, it gives input to function by passing parameter str
function string_reversed(str) {
  // return str.split("").reverse().join("");
  var reverse_str = "";
  for (var i = str.length - 1; i >= 0; i--) {
    reverse_str += str[i];
  }
  return reverse_str;
}

module.exports = string_reversed;
