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

function Maxsubarray(numsArr: number[]): number {
  let currSum = 0;
  let maxSum = -Infinity;

  //for-of Loop
  for (let num of numsArr) {
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
