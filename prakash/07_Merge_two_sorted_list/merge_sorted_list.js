//Function to merge two sorted linked lists
function MergetwoList(l1, l2) {
  var dummy = { val: 0, next: null };
  var current = dummy;
  while (l1 !== null && l2 !== null) {
    if (l1.val <= l2.val) {
      current.next = l1;
      l1 = l1.next;
    } else {
      current.next = l2;
      l2 = l2.next;
    }
    current = current.next;
  }
  current.next = l1 !== null && l1 !== void 0 ? l1 : l2;
  return dummy.next;
}
//Create Linked lists from arrays
function createLinkedList(arr) {
  if (arr.length === 0) return null;
  var dummy = { val: 0, next: null };
  var current = dummy;
  for (var i = 0; i < arr.length; i++) {
    current.next = { val: arr[i], next: null };
    current = current.next;
  }
  return dummy.next;
}
//Convert Linked lists to arrays
function linkedListToArray(head) {
  var arr = [];
  var current = head;
  for (; current !== null; current = current.next) {
    arr.push(current.val);
  }
  return arr;
}

module.exports = { MergetwoList, createLinkedList, linkedListToArray };
