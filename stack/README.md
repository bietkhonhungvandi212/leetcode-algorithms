# Matching pattern
## 1. Nested problem 
Stacks are naturally suited for parsing nested structures due to their LIFO property 
## 2. Monotoic stack
Using `Monotoic stack` maintains the elements in either increasing or decreasing order
- Next greater/ smaller element
->  Many algorithmic problems boil down to finding "the next/previous element that satisfies a certain condition" for each element in an array.

```golang
// Generic Monotonic Stack Template
func monotonicStackTemplate(nums []int) []int {
    stack := make([]int, 0)  // Usually store indices, not values
    result := make([]int, len(nums))
    
    for i, num := range nums {
        // Maintain monotonic property
        for len(stack) > 0 && shouldPop(stack[len(stack)-1], i, nums) {
            // The current element breaks monotonic property
            // This means we found the answer for the element being popped
            poppedIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            // Record the result - current element is the answer for popped element
            processResult(result, poppedIndex, i, nums)
        }
        
        stack = append(stack, i)
    }
    
    // Handle remaining elements in stack (usually no answer found)
    handleRemaining(stack, result)
    
    return result
}
```
## 3. Backtracking support
Stacks facilitate backtracking algorithms by maintaining decision points and allowing efficient state restoration