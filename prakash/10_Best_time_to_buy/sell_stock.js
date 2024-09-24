function maxProfit(prices) {
  var minPrice = Infinity;
  var maxProfit = 0;
  for (var _i = 0, prices_1 = prices; _i < prices_1.length; _i++) {
    var price = prices_1[_i];
    if (price < minPrice) {
      minPrice = price;
    }
    // Calculate the current profit and update maxProfit if it's larger than the previous max
    else if (price - minPrice > maxProfit) {
      maxProfit = price - minPrice;
    }
  }
  return maxProfit;
}
module.exports = maxProfit;
