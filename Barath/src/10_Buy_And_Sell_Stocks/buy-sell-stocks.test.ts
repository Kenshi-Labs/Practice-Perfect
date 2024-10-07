import { maxProfit } from "./buy-sell-stocks";

describe("maxProfit", () => {
    // 1. Test with an array where profit is possible
    test("should return maximum profit for a profitable array", () => {
        expect(maxProfit([7, 1, 5, 3, 6, 4])).toBe(5);
    });

    // 2. Test with an array where no profit is possible (descending order)
    test("should return 0 for a descending array", () => {
        expect(maxProfit([7, 6, 4, 3, 1])).toBe(0);
    });

    // 3. Test with an array containing duplicate values
    test("should return maximum profit for an array with duplicates", () => {
        expect(maxProfit([1, 2, 2, 3, 1, 4])).toBe(3);
    });

    // 4. Test with an array of length 2
    test("should return profit for an array of length 2", () => {
        expect(maxProfit([1, 2])).toBe(1);
    });
});