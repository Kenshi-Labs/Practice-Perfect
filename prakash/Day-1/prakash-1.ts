// Problem 1,

// TestCase1 - Test with a non-empty string
// Method-1 reverse array Method

const inputString : string = "hello";
const reversedString :string = reverseString(inputString);
console.log("Unit Test 1 =>" ,reversedString);

function reverseString(input: string):string {
    return input.split('').reverse().join("")
}

//
// Method-2 reverse loop method

const stringInput:string = "hello";    // Input
const str:string = stringsreverse(stringInput); //Input to function
console.log("Unit Test 1 =>",str)

function stringsreverse(input: string): string {
    let stringR: string = ""
for (let i=input.length - 1; i>=0; i--) {
     stringR +=input[i]
}
return stringR
}

// TestCase2 - Test with a empty string
//Method-1 reverse loop method

const empty: string = "";

const emptyS: string = emptySt(empty);
console.log("Unit Test 2 =>" ,emptyS==="")

function emptySt(empty: string): string {
    let test: string = "";
    for (let i = empty.length-1; i>=0; i--){
        test+=empty[i]
    }
    return test
}

// Method-2 

const emptyString : string = "";
const emptyStr :string = emptychar(emptyString);
console.log("Unit Test 2 =>" ,emptyStr==='');

function emptychar(str: string):string {
    return str.split('').reverse().join("")
}


// TestCase3 - Test with a string containing spaces
//Method-1 reverse array method

const spaces: string = "Hello World";
const Tspaces: string = TestSpaces(spaces);
const Tspace: string = checkSpaces(spaces);
console.log("Unit Test 3 =>" ,Tspaces);
console.log("Unit Test 3 =>" ,Tspace);

function TestSpaces(str: string): string{
    return str.split(/\b/).reverse().join('')
}
function checkSpaces(str: string): string{
    return str.split("").reverse().join('')
}

//Method-2 reverse loop method

const Space: string = "Hello World";
const tSpace = spacex(Space);
// const testwithspaces: string = testwithspace(Space);
console.log("Unit Test 3 =>" ,tSpace)
// console.log("Unit Test 3 =>" ,testwithspaces)

function spacex(str: string): string {
    let ilet: string = "";
    for (let i = str.length - 1; i>=0; i--){
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

const palindrome: string = "121";
const palindromestr: string = StrPalindrome(palindrome);
console.log("Unit Test 4 =>", palindromestr)

function StrPalindrome(str: string): string{
    let number: string = ""
    for(let i = str.length - 1; i>=0; i--){
        number +=str[i]
    }

    if (str === number){
        console.log("The String is Palindrom")
    } else {
        console.log("The String is Not Palindrom")
    }
    return number;
}


//Method-2 reverse array method

const palindrome1: string = "123";
const palindrome_no: string = funcpalindrome(palindrome1);
console.log("Unit Test 4 =>", palindrome_no)

function funcpalindrome(str: string): string {
    let a: string =  str.split("").reverse().join("")
    if (str === a){
        console.log("The String is Palindrom")
    } else {
        console.log("The String is Not Palindrom")
    }
    return a
}
