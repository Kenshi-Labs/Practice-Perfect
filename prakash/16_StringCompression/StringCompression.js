function CompressString(input) {
  if (input.length === 0) {
    return input;
  }
  var compressed = "";
  var count = 1;
  for (var i = 1; i < input.length; i++) {
    if (input[i] === input[i - 1]) {
      count++;
    } else {
      compressed += input[i - 1] + count;
      count = 1;
    }
  }
  compressed += input[input.length - 1] + count;
  return compressed.length < input.length ? compressed : input;
}
module.exports = CompressString;
