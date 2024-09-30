function climbStairs(n: number): number {
  if (n === 0 || n === 1) return 1;

  let first = 1;
  let second = 1;

  for (let i = 2; i <= n; i++) {
    const third = first + second;
    first = second;
    second = third;
  }

  return second;
}

export default climbStairs;
