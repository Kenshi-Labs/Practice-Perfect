export function compressString(input: any) {
  if (!input) return input; // If the string is empty, return it directly.

  let compressed = "";
  let count = 1;

  // Loop through the string and count repeated characters.
  for (let i = 1; i <= input.length; i++) {
    if (input[i] === input[i - 1]) {
      count++;
    } else {
      compressed += input[i - 1] + count;
      count = 1;
    }
  }

  // Return the original string if the compressed version is not smaller.
  return compressed.length < input.length ? compressed : input;
}
