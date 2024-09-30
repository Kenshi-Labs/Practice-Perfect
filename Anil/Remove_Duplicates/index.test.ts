import removeDuplicates from "./index";

test("array containing duplicates", () => {
  const nums: (number | string)[] = [1, 1, 2];
  const length = removeDuplicates(nums);
  expect(length).toBe(2);
  expect(nums).toEqual([1, 2, "_"]);
});

test("array with no duplicates", () => {
  const nums: (number | string)[] = [1, 2, 3];
  const length = removeDuplicates(nums);
  expect(length).toBe(3);
  expect(nums).toEqual([1, 2, 3]);
});

test("empty array", () => {
  const nums: (number | string)[] = [];
  const length = removeDuplicates(nums);
  expect(length).toBe(0);
  expect(nums).toEqual([]);
});

test("array containing all identical elements", () => {
  const nums: (number | string)[] = [1, 1, 1, 1];
  const length = removeDuplicates(nums);
  expect(length).toBe(1);
  expect(nums).toEqual([1, "_", "_", "_"]);
});
