// function MaxsubArray(numsArr: number[]):number {
//   let currSum = 0;
//   let maxSum = -Infinity;
//for loop
//   for (let i = 0; i < numsArr.length; i++) {
// currSum = Math.max(numsArr[i], currSum + numsArr[i]);
// maxSum = Math.max(currSum, maxSum);
// if (currSum + numsArr[i] > numsArr[i]) {
//     currSum = currSum + numsArr[i]
// } else {
//     currSum = numsArr[i];
// }
// if(currSum > numsArr[i]){
//     maxSum = currSum
// }
//   }
//   return maxSum;
// }
function MaxsubArray(numsArr) {
  var currSum = 0;
  var maxSum = -Infinity;
  //for-of Loop
  for (var _i = 0, numsArr_1 = numsArr; _i < numsArr_1.length; _i++) {
    var num = numsArr_1[_i];
    if (currSum + num > num) {
      currSum = currSum + num;
    } else {
      currSum = num;
    }
    if (currSum > maxSum) {
      maxSum = currSum;
    }
  }
  return maxSum;
}

module.exports = MaxsubArray;
