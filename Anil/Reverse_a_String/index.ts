function reverseString(str: string) {
  let reversedStr = "";
  for (let i = str.length - 1; i >= 0; i--) {
    reversedStr += str[i];
  }
  return reversedStr;
}

const strr = reverseString("hello");
console.log(strr);

export default reverseString;
