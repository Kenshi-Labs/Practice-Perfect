const maxProfit = require("./sell_stock");

describe("maxProfit", () => {
  test("Test with an array where profit is possible", () => {
    const prices = [7, 1, 5, 3, 6, 4];
    const output = 5;
    expect(maxProfit(prices)).toBe(output);
  });

  test("Test with an array where no profit is possible (descending order)", () => {
    const prices = [7, 6, 4, 3, 1];
    const output = 0;
    expect(maxProfit(prices)).toBe(output);
  });

  test("should handle arrays with duplicate values", () => {
    const prices = [2, 4, 1, 4, 4, 5, 3, 5, 1];
    const output = 4; // Buy at 1, sell at 5
    expect(maxProfit(prices)).toBe(output);
  });

  test("should return the profit or 0 for an array of length 2", () => {
    const prices1 = [3, 5];
    const prices2 = [5, 3];
    expect(maxProfit(prices1)).toBe(2); // Profit = 2
    expect(maxProfit(prices2)).toBe(0); // No profit
  });
});
