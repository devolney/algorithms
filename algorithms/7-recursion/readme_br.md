# [BR] Recursividade Entendendo o Conceito e Aplicações

## 1. Introdução

A recursividade é uma técnica de programação poderosa e elegante que permite que uma função chame a si mesma para resolver problemas. Essa abordagem é particularmente útil para problemas que podem ser divididos em subproblemas menores e resolvidos de forma repetitiva.

Neste artigo, exploraremos o que é recursividade, suas vantagens, desvantagens, o conceito de **Divide and Conquer** e um exemplo prático usando a sequência de Fibonacci. Também veremos como otimizar a recursividade com [memoization](https://www.geeksforgeeks.org/what-is-memoization-a-complete-tutorial/).

---

## 2. O que é Recursividade?

Recursividade ocorre quando uma função realiza chamadas a si mesma para resolver uma parte do problema maior. Para que funcione corretamente, duas coisas são essenciais:

1. **Caso Base**: A condição que encerra a recursão.
2. **Chamada Recursiva**: O momento em que a função se chama novamente, aproximando-se do caso base.

---

## 3. Estratégia "Divide and Conquer"

Recursividade é frequentemente associada à estratégia **Divide and Conquer**, que funciona assim:

1. Divida o problema em subproblemas menores.
2. Resolva cada subproblema recursivamente.
3. Combine as soluções para obter o resultado final.

Exemplos clássicos que utilizam essa abordagem incluem algoritmos como Merge Sort, Quick Sort e pesquisa binária.

---

## 4. Exemplos Práticos: Sequência de Fibonacci

A sequência de Fibonacci é um excelente exemplo para entender recursividade. A sequência é definida como:

- $F(0) = 0, F(1) = 1$
- Para $n > 1: F(n) = F(n-1) + F(n-2)$

Abaixo, exploramos três abordagens para calcular Fibonacci.

### 4.1. Fibonacci Sem Recursividade (Iterativo)

Aqui utilizamos um loop para calcular Fibonacci de forma eficiente.

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

### 4.2. Fibonacci Com Recursividade

Uma abordagem simples, mas não otimizada, onde cada chamada recursiva resolve dois subproblemas.

```go
func fibonacciRecursive(n int) int {
    if n <= 1 {
    return n
    }
    return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
```

- **Desvantagem:** A mesma função é chamada repetidamente para os mesmos valores, resultando em cálculos redundantes.

### 4.3. Fibonacci Com Recursividade e Memoization

Memoization é uma técnica de otimização que armazena os resultados das funções já calculadas, evitando chamadas repetidas.

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

---

## 5. Vantagens e Desvantagens da Recursividade

### 5.1 Vantagens

- **Clareza e Elegância:** Torna o código mais intuitivo para problemas complexos.
- **Adequado para estruturas naturais:** Como árvores, grafos e algoritmos de divisão.

### 5.2 Desvantagens

- **Desempenho:** Cada chamada recursiva ocupa espaço na pilha de chamadas, podendo causar **stack overflow**.
- **Sobreposição de cálculos:** Sem otimização, cálculos redundantes podem ocorrer.

## 6. Comparação de Desempenho

| Abordagem                | Complexidade Temporal | Complexidade Espacial |
|--------------------------|-----------------------|-----------------------|
| Iterativa                | O(n)                  | O(1)                  |
| Recursiva Simples        | O(2^n)                | O(n)                  |
| Recursiva com Memoization| O(n)                  | O(n)                  |

## 7. Conclusão

A recursividade é uma ferramenta poderosa, mas deve ser usada com cuidado. Ela é ideal para problemas naturalmente recursivos, como árvores e grafos, mas pode ser ineficiente sem otimização. A técnica de memoization ajuda a melhorar o desempenho, tornando a recursividade uma abordagem prática para muitos casos.

## 8. Instagram

Na postagem do Instagram [Algoritmos - O que é recursividade?](https://www.instagram.com/p/DFLuE0jylUW/?img_index=1) explico o que é resursividade quais são as suas vantagens e como utilizar de forma otimizada.
