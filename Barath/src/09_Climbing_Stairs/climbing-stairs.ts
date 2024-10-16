
function climbStairs(n: number): number {
    if (n <= 2) return n;
    let prev = 1;
    let curr = 2;
    for (let i = 3; i <= n; i++) {
        const temp = curr;
        curr = prev + curr;
        prev = temp;
    }
    return curr;
}

console.log(climbStairs(2));
console.log(climbStairs(3));

export { climbStairs };

