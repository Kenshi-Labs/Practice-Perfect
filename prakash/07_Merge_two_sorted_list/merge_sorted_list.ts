//Define the ListNode type
type ListNode = {
  val: number;
  next: ListNode | null;
};

//Function to merge two sorted linked lists

function MergetwoList(
  l1: ListNode | null,
  l2: ListNode | null
): ListNode | null {
  const dummy: ListNode = { val: 0, next: null };
  let current = dummy;

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

  current.next = l1 ?? l2;
  return dummy.next;
}

//Create Linked lists from arrays
function createLinkedList(arr: number[]): ListNode | null {
  if (arr.length === 0) return null;
  const dummy: ListNode = { val: 0, next: null };

  let current = dummy;

  for (let i = 0; i < arr.length; i++) {
    current.next = { val: arr[i], next: null };
    current = current.next;
  }
  return dummy.next;
}

//Convert Linked lists to arrays
function linkedListToArray(head: ListNode | null): number[] {
  const arr: number[] = [];
  let current = head;

  for (; current !== null; current = current.next) {
    arr.push(current.val);
  }

  return arr;
}
