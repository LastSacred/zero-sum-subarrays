You are given an array of integers. You need to print the total count of sub-arrays having their sum equal to 0

Input:
First line of the input contains an integer T denoting the number of test cases. Each test case consists of two lines. First line of each test case contains an Integer N denoting the size of the array, and the second line contains N space separated integers.

Output:
For each test case, in a new line, print the total number of subarrays whose sum is equal to 0.

Expected Time Complexity: O(N).
Expected Auxiliary Space: O(N).

Constraints:    
1 <= T <= 100
1 <= N <= 107
-1010 <= Ai <= 1010

Example:
Input:
2
6
0 0 5 5 0 0
10
6  -1 -3 4 -2 2 4 6 -12 -7
Output:
6
4

Explanation:
Testcase 1: There are 6 subarrays present whose sum is zero. The starting and ending indices are:
(0,0) (1,1) (0,1) (4,4) (5,5) (4,5)
Testcase 2: There are 4 subarrays present whose sum is zero. The starting and ending indices are:
(1,3) (4,5) (1,5) (5,8)