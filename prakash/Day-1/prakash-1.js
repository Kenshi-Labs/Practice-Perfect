// Problem 1,
// TestCase1 - Test with a non-empty string
// Method-1 reverse array Method
var inputString = "hello";
var reversedString = reverseString(inputString);
console.log("Unit Test 1 =>", reversedString);
function reverseString(input) {
    return input.split('').reverse().join("");
}
//
// Method-2 reverse loop method
var stringInput = "hello"; // Input
var str = stringsreverse(stringInput); //Input to function
console.log("Unit Test 1 =>", str);
function stringsreverse(input) {
    var stringR = "";
    for (var i = input.length - 1; i >= 0; i--) {
        stringR += input[i];
    }
    return stringR;
}
// TestCase2 - Test with a empty string
//Method-1 reverse loop method
var empty = "";
var emptyS = emptySt(empty);
console.log("Unit Test 2 =>", emptyS === "");
function emptySt(empty) {
    var test = "";
    for (var i = empty.length - 1; i >= 0; i--) {
        test += empty[i];
    }
    return test;
}
// Method-2 
var emptyString = "";
var emptyStr = emptychar(emptyString);
console.log("Unit Test 2 =>", emptyStr === '');
function emptychar(str) {
    return str.split('').reverse().join("");
}
// TestCase3 - Test with a string containing spaces
//Method-1 reverse array method
var spaces = "Hello World";
var Tspaces = TestSpaces(spaces);
var Tspace = checkSpaces(spaces);
console.log("Unit Test 3 =>", Tspaces);
console.log("Unit Test 3 =>", Tspace);
function TestSpaces(str) {
    return str.split(/\b/).reverse().join('');
}
function checkSpaces(str) {
    return str.split("").reverse().join('');
}
//Method-2 reverse loop method
var Space = "Hello World";
var tSpace = spacex(Space);
// const testwithspaces: string = testwithspace(Space);
console.log("Unit Test 3 =>", tSpace);
// console.log("Unit Test 3 =>" ,testwithspaces)
function spacex(str) {
    var ilet = "";
    for (var i = str.length - 1; i >= 0; i--) {
        ilet += str[i];
    }
    return ilet;
}
// function testwithspace(str: string): string {
//     let Stest =  "";
//     let word = "";
// for (let i = str.length - 1; i>=0; i++){
//     Stest = word + " " + Stest;
// }
// return Stest;
// }
// TestCase4 - Test with a palindrome
//Method-1 reverse loop method
var palindrome = "121";
var palindromestr = StrPalindrome(palindrome);
console.log("Unit Test 4 =>", palindromestr);
function StrPalindrome(str) {
    var number = "";
    for (var i = str.length - 1; i >= 0; i--) {
        number += str[i];
    }
    if (str === number) {
        console.log("The String is Palindrom");
    }
    else {
        console.log("The String is Not Palindrom");
    }
    return number;
}
//Method-2 reverse array method
var palindrome1 = "123";
var palindrome_no = funcpalindrome(palindrome1);
console.log("Unit Test 4 =>", palindrome_no);
function funcpalindrome(str) {
    var a = str.split("").reverse().join("");
    if (str === a) {
        console.log("The String is Palindrom");
    }
    else {
        console.log("The String is Not Palindrom");
    }
    return a;
}
