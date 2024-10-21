"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const index_1 = __importDefault(require("./index"));
test("Test with valid anagrams", () => {
    expect((0, index_1.default)("anagram", "nagaram")).toBe(true);
});
test("Test with strings of different lengths", () => {
    expect((0, index_1.default)("anil", "sunil")).toBe(false);
});
test("Test with strings containing the same letters but different frequencies", () => {
    expect((0, index_1.default)("aabbcc", "aabbc")).toBe(false);
});
test("Test with empty strings", () => {
    expect((0, index_1.default)("", "")).toBe(true);
});
