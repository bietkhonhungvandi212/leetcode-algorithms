1. Regconition
- This is a pattern that sastify DP condition: 
    + To solve a problem, you must solve its subproblems in a similar way. And there is always a optimal solution to solve all the parent problem and sub problem in recursive way. 
    + We can create a common formula for addressing this problem. If we have a length from position `start` to `end`, every cut position in array `cuts` in that range, we have
        + T(start, end) = Min {T(start, k) + T(k, end) + length(end - start)}
    + Every subproblem is dependent from each other, it means if the cost T(start - end) is min, so every cost of subproblems which were cut must be min cost, eventhough they are subproblems to solve the parent problem 