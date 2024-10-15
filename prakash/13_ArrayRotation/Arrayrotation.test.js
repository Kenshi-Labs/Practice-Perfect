const rotateArray = require("./Arrayrotation");
describe("rotateArray", () => {
  test("Test with a non-empty array and positive rotation", () => {
    const nums = [1, 2, 3, 4, 5];
    rotateArray(nums, 2);
    expect(nums).toEqual([4, 5, 1, 2, 3]);
  });

  test("Test with rotation equal to array length", () => {
    const nums = [1, 2, 3, 4, 5];
    rotateArray(nums, 5);
    expect(nums).toEqual([1, 2, 3, 4, 5]);
  });

  test("Test with rotation greater than array length", () => {
    const nums = [1, 2, 3, 4, 5];
    rotateArray(nums, 7); // Equivalent to rotating by 2 steps
    expect(nums).toEqual([4, 5, 1, 2, 3]);
  });

  test("Test with an empty array", () => {
    const nums = [];
    rotateArray(nums, 3);
    expect(nums).toEqual([]);
  });
});
