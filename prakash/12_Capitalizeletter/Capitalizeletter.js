function CapitalizefirstLetter(input) {
  var result = "";
  var capitalizeNext = true;
  for (var i = 0; i < input.length; i++) {
    var char = input[i];
    if (char === "" || char.match(/[^a-zA-Z0-9]/)) {
      result += char;
      capitalizeNext = true;
    } else if (capitalizeNext && char >= "a" && char <= "z") {
      result += String.fromCharCode(char.charCodeAt(0) - 32);
      capitalizeNext = false;
    } else {
      result += char;
      capitalizeNext = false;
    }
  }
  return result;
}

module.exports = CapitalizefirstLetter;
