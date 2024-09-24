function maxProfit(prices: number[]): number {
  let minPrice = Infinity;
  let maxProfit = 0;

  for (let price of prices) {
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
