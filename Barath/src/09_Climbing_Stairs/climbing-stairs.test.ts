import { climbStairs } from "./climbing-stairs";

describe("climbStairs", () => {
    it("should return 1 for n = 1", () => {
        expect(climbStairs(1)).toBe(1);
    });

    it("should return 2 for n = 2", () => {
        expect(climbStairs(2)).toBe(2);
    });

    it("should return 8 for n = 5", () => {
        expect(climbStairs(5)).toBe(8);
    });

    it("should return 0 for n = 0", () => {
        expect(climbStairs(0)).toBe(0);
    });
});
