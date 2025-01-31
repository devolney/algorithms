# [EN] Recursion Understanding the Concept and Applications

## 1. Introduction

Recursion is a powerful and elegant programming technique that allows a function to call itself to solve problems. This approach is particularly useful for problems that can be divided into smaller subproblems and solved repeatedly.

In this article, we will explore what recursion is, its advantages, disadvantages, the concept of **Divide and Conquer** and a practical example using the Fibonacci sequence. We will also see how to optimize recursion with [memoization](https://www.geeksforgeeks.org/what-is-memoization-a-complete-tutorial/).

---

## 2. What is Recursion?

Recursion occurs when a function calls itself to solve a part of the larger problem. For it to work correctly, two things are essential:

1. **Base Case**: The condition that ends the recursion.

2. **Recursive Call**: The moment when the function calls itself again, getting closer to the base case.

---

## 3. "Divide and Conquer" Strategy

Recursion is often associated with the **Divide and Conquer** strategy, which works like this:

1. Divide the problem into smaller subproblems.

2. Solve each subproblem recursively.

3. Combine the solutions to obtain the final result.

Classic examples that use this approach include algorithms such as Merge Sort, Quick Sort and binary search.

---

## 4. Practical Examples: Fibonacci Sequence

The Fibonacci sequence is an excellent example to understand recursion. The sequence is defined as:

- $F(0) = 0, F(1) = 1$
- For $n > 1: F(n) = F(n-1) + F(n-2)$

Below, we explore three approaches to calculating Fibonacci.

### 4.1. Fibonacci Without Recursion (Iterative)

Here we use a loop to calculate Fibonacci efficiently.

```go
func fibonacciIterative(n int) int {
    if n <= 1 {
    return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
    a, b = b, a+b
    }
    return b
}
```

### 4.2. Fibonacci With Recursion

A simple but unoptimized approach, where each recursive call solves two subproblems.

```go
func fibonacciRecursive(n int) int {
    if n <= 1 {
    return n
    }
    return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
```

- **Disadvantage:** The same function is called repeatedly for the same values, resulting in redundant calculations.

### 4.3. Fibonacci With Recursion and Memoization

Memoization is an optimization technique that stores the results of already calculated functions, avoiding repeated calls.

```go
var memo = make(map[int]int)

func fibonacciMemoization(n int) int {
    if n <= 1 {
    return n
    }
    if _, exists := memo[n]; !exists {
    memo[n] = fibonacciMemoization(n-1) + fibonacciMemoization(n-2)
    }
    return memo[n]
}
```

### 4.4. Execution time comparison

Run the shell command below to see how long it takes to execute the algorithm in each scenario presented. Feel free to change the `n` variable in [code](main.go) to see how the algorithm performs.

```sh
make recursion
```

---

## 5. Advantages and Disadvantages of Recursion

### 5.1 Advantages

- **Clarity and Elegance:** Makes code more intuitive for complex problems.

- **Suitable for natural structures:** Such as trees, graphs, and division algorithms.

### 5.2 Disadvantages

- **Performance:** Each recursive call takes up space in the call stack, which can cause **stack overflow**.
- **Overlapping of calculations:** Without optimization, redundant calculations can occur.

## 6. Performance Comparison

| Approach                  | Time Complexity       | Space Complexity      |
|---------------------------|-----------------------|-----------------------|
| Iterative                 | O(n)                  | O(1)                  |
| Simple Recursive          | O(2^n)                | O(n)                  |
| Recursive with Memoization| O(n)                  | O(n)                  |

## 7. Conclusion

Recursion is a powerful tool, but it should be used with caution. It is ideal for naturally recursive problems, such as trees and graphs, but it can be inefficient without optimization. The memoization technique helps improve performance, making recursion a practical approach for many cases.

---

## 8. Instagram

In the Instagram post [Algoritmos - O que é recursividade?](https://www.instagram.com/p/DFLuE0jylUW/?img_index=1) I explain what recursion is, what its advantages are and how to use it optimally.
