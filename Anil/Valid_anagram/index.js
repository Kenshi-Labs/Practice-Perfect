"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function anagram(str1, str2) {
    if (str1.length !== str2.length) {
        return false;
    }
    let arr = str1.split("").sort().join("");
    let arr1 = str2.split("").sort().join("");
    return arr === arr1;
}
let str = "anagram";
let str1 = "nagrama";
console.log(anagram(str, str1));
exports.default = anagram;
