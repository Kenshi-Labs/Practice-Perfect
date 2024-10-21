Small Story :)

I have two boxes of toys, and each toy has a number written on it. One box has toys numbered 1, 2, 4, and the other box has toys numbered 1, 3, 4. The toys in both boxes are already neatly arranged in order, from the smallest to the biggest.

Now, my job is to make one big line of toys by picking toys from both boxes, but I need to keep the order from smallest to biggest.

Here’s the plan:

I start with two boxes of toys. Let's call them Box 1 and Box 2.
I have an empty shelf called "Result" where I'm going to place the toys in order.

I look at the first toy in Box 1 and Box 2. Which one is smaller?
If the toy from Box 1 is smaller or equal to the toy from Box 2, I put the toy from Box 1 on the shelf. Then I move to the next toy in Box 1.
If the toy from Box 2 is smaller, I place that toy on the shelf and move to the next toy in Box 2.
I keep doing this until one of the boxes runs out of toys.
But wait! What if one of the boxes still has toys left? Don’t worry—once one box is empty, I can take all the remaining toys from the other box and add them to the shelf.

In the code:

I compare the toys (numbers in the lists) one by one using a loop.
If the toy from list1 (Box 1) is smaller, I add it to result and move to the next toy in list1.
If the toy from list2 (Box 2) is smaller, I add it to result and move to the next toy in list2.
When I finish looking through one box, I add all the toys from the other box.
The result is one big, sorted line of toys on my shelf!

In the example of [1, 2, 4] and [1, 3, 4], the final line would be: 1, 1, 2, 3, 4, 4.