// n=1 => 1;   No. of ways to climb stairs; (1)
// n=2 => 2;   No. of ways to climb stairs; (2), (11)
// n=3 => 3;   No. of ways to climb stairs; (111), (21), (12)
// n=4 => 5;   No. of ways to climb stairs; (1111), (112), (211), (121), (22)
// n=5 => 8;   No. of ways to climb stairs; (1111), (221), (212), (1121), (122), (2111), (1211), (1112)

//f(n) = f(n-1) + f(n-2)
function climbing_stairs(steps) {
  if (steps <= 2) {
    return steps;
  }
  var dp = new Array(steps + 1);
  dp[1] = 1;
  dp[2] = 2;
  for (var i = 3; i <= steps; i++) {
    dp[i] = dp[i - 1] + dp[i - 2];
  }
  return dp[steps];
}

module.exports = climbing_stairs;
