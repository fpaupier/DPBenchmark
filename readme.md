# Benchmark dynamic programming en Go

Le _Dynamic Programming_ permet d'avoir une implémentation optimisée en temps d'exécution par rapport
à une implémentation récursive. 

L'exemple classique pour illustrer ce gain est celui de la suite de Fibonacci pour laquelle on peut passer
d'une complexité d'`O(2ˆn)` à `O(n)` en temps d'exécution.

> En mathématiques, la suite de Fibonacci est une suite d'entiers dans laquelle chaque terme est la somme des deux termes qui le précèdent. Elle commence par les termes 0 et 1
>
>   La suite est définie par F0 = 0 , F1 = 1 , et par Fn=F(n-1)+F(n-2) pour n > 1.
>
> [Wikipedia](https://fr.wikipedia.org/wiki/Suite_de_Fibonacci)

## Implémentation récursive
Une implémentation récursive est intuitive et on peut proposer en Go:

````go
package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}
````

On remarque que chaque appel à ``fib`` déclenche l'appel à deux fonctions ``fib``. Exemple pour ``fib(5)``ci-dessous.
Le terme `fib(2)` est calculé deux fois.

_Figure 1. Arbre de résolution du 5ᵉ terme de la suite de Fibonacci

![Arbre de résolution du terme 5 de la suite de Fibonacci](assets/fib(5).jpeg)

## Implémentation itérative
Une première optimisation consiste à utiliser une approche itérative en conservant, dans une liste, les éléments déjà calculés.
On complète ce tableau - nommé `dp` pour `dynamic programming` - en commençant par les conditions initiales (les termes
`fib(0)` et `fib(1)` d'initialisation de la récurrence) et en remontant jusqu'au terme `n` souhaité. C'est l'approche _bottom-up_.

````go
package main

func fibIter(n int) int {
	dp := make ([]int, n + 1 , n + 1)
	dp[1] = 1
	for i := 2; i < n + 1; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}
````

## Avec memoization
Une dernière optimisation permet de réduire l'usage en mémoire - `O(n)` dans cette version itérative - à `O(1)`.
Le n-ème terme en dépend que du n-1ème et du n-2ème, on peut donc éliminer l'usage d'un tableau et se contenter de trois
variables. 
 - `cur` pour le terme actuel de la suite (`fib(i)`),
 - `prev` pour le terme précédent (`fib(i-1)`)
 - `pprev` pour le terme encore précédent (`fib(i-2)`)

````go
package main

func fibOpti(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	pprev := 0
	prev := 1
	var cur int
	for i := 2; i < n + 1; i++ {
		cur = prev + pprev
		pprev = prev
		prev = cur
	}
	return cur
}
````
On a une empreinte mémoire en `O(1)` et un temps d'exécution en `O(n)` avec la `for-loop`.


### Références

- [Cracking the Coding Interview](https://www.crackingthecodinginterview.com/), chapitre 8, Dynamic Programming

- Pour des exercices sur le DP, voir la catégorie [Dynamic Programing](https://leetcode.com/problemset/all/?topicSlugs=dynamic-programming) sur leetcode. 
  Le [#53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) est un classique.
  
- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go), de Dave Cheney