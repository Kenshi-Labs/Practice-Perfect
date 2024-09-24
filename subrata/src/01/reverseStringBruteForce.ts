function reverseStringBruteForce(input: string): string {
  let reversed = "";
  // perform the calculation
  for (let i = input.length - 1; i >= 0; i--) {
    reversed += input[i];
  }
  return reversed;
}

export default reverseStringBruteForce;
