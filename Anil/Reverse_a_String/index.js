"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function reverseString(str) {
    let reversedStr = "";
    for (let i = str.length - 1; i >= 0; i--) {
        reversedStr += str[i];
    }
    return reversedStr;
}
const strr = reverseString("hello");
console.log(strr);
exports.default = reverseString;
