let a = 2.123;
const result = math(a);
console.log(result);

// function math(b) {
//   return (b | 0) + (b % 1 !== 0 ? 1 : 0);
// }

function math(b) {
  let strr = b.toString();

  let decimalIndex = strr.indexOf(".");

  if (decimalIndex === -1) {
    return b;
  }

  let integer = strr.slice(0, decimalIndex);
  let decimal = strr.slice(decimalIndex + 1);
  console.log(integer);
  console.log(decimal);

  if (decimal !== 0) {
    integer = (parseInt(integer) + 1).toString();
  }
  return parseInt(integer);
}
