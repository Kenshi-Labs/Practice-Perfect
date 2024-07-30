function mergeSortedLists(arr1: number[], arr2: number[]): number[] {
    let merged: number[] = [];
    let i = 0, j = 0;

    while (i < arr1.length && j < arr2.length) {
        if (arr1[i] < arr2[j]) {
            merged.push(arr1[i++]);
        } else {
            merged.push(arr2[j++]);
        }
    }

    return [...merged, ...arr1.slice(i), ...arr2.slice(j)];
}

let l1 = [1,2,4]
let l2 = [1,3,4]

console.log(mergeSortedLists(l1, l2));


export { mergeSortedLists };
