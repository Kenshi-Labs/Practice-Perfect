"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const index_1 = __importDefault(require("./index"));
test("Test with a non-empty string", () => {
    expect((0, index_1.default)("hello")).toBe("olleh");
});
test("Test with an empty string", () => {
    expect((0, index_1.default)("")).toBe("");
});
test("Test with a string containing spaces", () => {
    expect((0, index_1.default)("hello world")).toBe("dlrow olleh");
});
test("Test with a palindrome", () => {
    expect((0, index_1.default)("dad")).toBe("dad");
});
