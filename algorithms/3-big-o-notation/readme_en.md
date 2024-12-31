# [EN] Big O Notation

Big O Notation is a mathematical notation used to describe the performance or complexity of an algorithm. It provides a way to express how the runtime or space requirements of an algorithm grow relative to the input size. In this guide, we will explore Big O Notation in detail, its importance, and how to calculate it, with examples and visual illustrations.

---

## 1. What is Big O Notation?

Big O Notation is used to classify algorithms based on their growth rates. It focuses on the worst-case scenario, which helps developers understand the upper bound of an algorithm's performance.

**Key Concept:**

- **Input size ($n$):** Represents the size of the input data.
- **Growth rate:** Describes how the runtime or space grows as input size increases.
- **Dominant term:** The term in the function that grows fastest and defines the complexity.

> Example: If an algorithm takes $2n^2 + 3n + 1$ operations, Big O focuses only on $n^2$, as it dominates the growth.

---

## 2. Why is Big O Notation Important?

### 2.1. Scalability Analysis

Big O helps predict how an algorithm scales with larger input sizes. Algorithms that work well for small inputs may perform poorly with large datasets.

### 2.2. Performance Comparison

It allows developers to compare different algorithms and choose the most efficient one for a given problem.

### 2.3. Optimization

Identifying inefficiencies through Big O enables optimizations to improve performance.

---

## 3. Common Big O Complexities

### 3.1. Constant Time - $O(1)$

The runtime does not depend on the input size.

```go
func getFirstElement(arr []int) int {
    return arr[0] // Always 1 operation.
}
```

> Example: Accessing an element in an array.

### 3.2. Linear Time - $O(n)$

The runtime grows linearly with the input size.

```go
func printElements(arr []int) {
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i]) // Executes n times.
    }
}
```

> Example: Iterating through an array.

### 3.3. Quadratic Time - $O(n^2)$

The runtime grows proportionally to the square of the input size.

```go
func checkPairs(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] == arr[j] {
                fmt.Println("Match")
            }
        }
    }
}
```

> Example: Nested loops for pair comparisons.

### 3.4. Logarithmic Time - $O(log n)$

The runtime decreases as the input size grows, often appearing in divide-and-conquer algorithms.

```go
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    for left <= right {
        mid := (left + right) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

> Example: Binary search algorithm.

### 3.5. Exponential Time - $O(2^n)$

The runtime doubles with each additional input, common in recursive algorithms.

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

> Example: Calculating Fibonacci numbers.

### 3.6. Grafical Analisys

The graph below shows the growth curves of all the complexities discussed above, illustrating their differences as the input size increases.

![cost function curves](graph.png)

---

## 4. How to Calculate Big O Complexity

### 4.1. Step 1: Analyze Each Operation

Go through the code line-by-line and count operations.

### 4.2. Step 2: Remove Constants

Ignore constant terms and coefficients, as they become negligible for large inputs.

### 4.3. Step 3: Focus on Dominant Terms

Keep only the term with the highest growth rate.

#### Example Analysis

**Algorithm:**

```go
func example(arr []int) {
    sum := 0           // 1 operation (O(1))
    for i := 0; i < len(arr); i++ { // O(n)
        for j := 0; j < len(arr); j++ { // O(n)
            sum += arr[j]   // O(1)
        }
    }
    fmt.Println(sum)   // O(1)
}
```

**Step 1:** Count operations:

- Constant: $O(1)$ for initialization.
- Outer loop: $O(n)$.
- Inner loop: $O(n)$.
- Total: $O(1 + n + n^2 + 1)$.

**Step 2:** Remove constants:

- $O(n^2)$

**Step 3:** Final complexity:

- $O(n^2)$ (Quadratic).

---

### 5. Key Takeaways

1. Big O Notation evaluates how algorithms scale with input size.
2. Focus on dominant terms to determine complexity.
3. Use it to compare algorithms and optimize performance.
4. Visualize growth curves to better understand scalability.

---

### 6. Conclusion

Big O Notation is a fundamental concept for analyzing and improving algorithm performance. By understanding its principles and applying them to real-world scenarios, developers can make informed decisions to build scalable and efficient applications.

---

### 7. Instagram

In the Instagram post [Algoritmos - Big O Notation](https://www.instagram.com/p/DDw6DF1O9yP/?img_index=1) I explain what is Big O Notation, why it is important and how to calculate it.
