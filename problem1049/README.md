## Solution 1

We are add/subtract each number from the array, so we can enumerate all the possible sums and find the least one.
For each array of numbers, we can get the possible sums from its sub-array by add/subtract the last number from sub-array's each possible sum.

Cut the branch:
We only have to mark the absolute value of sums, as there will always be an opposite value for the sum.

sum[i]:each = sum[i-1]:each +/- stone[i]

## Solution 2

0-1 bag
We can divide the stones into 2 groups, one with positive sign, and the other with negative sign. To find the least possible stone weight, the negative group stones must have a sum less than but close to the sum of all stones.

Here we have a bag of volumn floor(sum/2), and we want to fill it as much as possible.