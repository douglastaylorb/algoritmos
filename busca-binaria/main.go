package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== TESTES BÁSICOS ===")
	basicTests()

	fmt.Println("\n=== PODER DA BUSCA BINÁRIA ===")
	powerDemo()

	fmt.Println("\n=== COMPARAÇÃO DE PERFORMANCE ===")
	performanceComparison()

	fmt.Println("\n=== VISUALIZAÇÃO PASSO A PASSO ===")
	stepByStepDemo()

	fmt.Println("\n=== TESTE FUNÇÃO RECURSIVA ===")
	recursiveTests()

	fmt.Println("\n=== COMPARAÇÃO: ITERATIVA vs RECURSIVA ===")
	iterativeVsRecursive()
}

func basicTests() {
	// Array ordenado para os testes
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Casos de teste com valor esperado
	testCases := []struct {
		value    int // Valor a ser buscado
		expected int // Índice esperado (-1 se não encontrado)
	}{
		{1, 0},   // Primeiro elemento
		{10, 9},  // Último elemento
		{5, 4},   // Elemento do meio
		{7, 6},   // Elemento qualquer
		{11, -1}, // Não existe (maior que todos)
		{0, -1},  // Não existe (menor que todos)
		{-5, -1}, // Não existe (muito menor)
	}

	// Executa cada caso de teste
	for _, test := range testCases {
		result := checkBinary(list, test.value)
		status := "✅"
		if result != test.expected {
			status = "❌"
		}
		fmt.Printf("%s Buscando %d: encontrado no índice %d (esperado: %d)\n",
			status, test.value, result, test.expected)
	}
}

// ✅ NOVA FUNÇÃO: Testa a versão recursiva
func recursiveTests() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	testCases := []struct {
		value    int
		expected int
	}{
		{1, 0},   // Primeiro elemento
		{10, 9},  // Último elemento
		{5, 4},   // Elemento do meio
		{7, 6},   // Elemento qualquer
		{11, -1}, // Não existe (maior)
		{0, -1},  // Não existe (menor)
	}

	for _, test := range testCases {
		result := recursiveCheckBinary(list, test.value, len(list)-1, 0)
		status := "✅"
		if result != test.expected {
			status = "❌"
		}
		fmt.Printf("%s [RECURSIVA] Buscando %d: encontrado no índice %d (esperado: %d)\n",
			status, test.value, result, test.expected)
	}
}

// Compara performance iterativa vs recursiva
func iterativeVsRecursive() {
	size := 100000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i * 2
	}

	target := 50000 // Elemento a ser buscado

	start := time.Now()
	iterativeResult := checkBinary(arr, target)
	iterativeTime := time.Since(start)

	start = time.Now()
	recursiveResult := recursiveCheckBinary(arr, target, len(arr)-1, 0)
	recursiveTime := time.Since(start)

	fmt.Printf("Versão Iterativa: índice %d em %v\n", iterativeResult, iterativeTime)
	fmt.Printf("Versão Recursiva: índice %d em %v\n", recursiveResult, recursiveTime)

	if iterativeTime < recursiveTime {
		ratio := float64(recursiveTime.Nanoseconds()) / float64(iterativeTime.Nanoseconds())
		fmt.Printf("Iterativa foi %.2fx mais rápida\n", ratio)
	} else {
		ratio := float64(iterativeTime.Nanoseconds()) / float64(recursiveTime.Nanoseconds())
		fmt.Printf("Recursiva foi %.2fx mais rápida\n", ratio)
	}
}

func powerDemo() {
	sizes := []int{100, 1000, 10000, 100000, 1000000}

	for _, size := range sizes {
		// Criando array ordenado de números pares
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = i * 2
		}

		// Buscando um elemento próximo ao meio
		target := size - 2

		// Medindo tempo de execução
		start := time.Now()
		index := checkBinary(arr, target)
		duration := time.Since(start)

		fmt.Printf("Array de %d elementos: encontrou %d no índice %d em %v\n",
			size, target, index, duration)
	}
}

func performanceComparison() {
	size := 1000000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i // Números sequenciais
	}

	target := 750000

	start := time.Now()
	linearIndex := linearSearch(arr, target)
	linearTime := time.Since(start)

	start = time.Now()
	binaryIndex := checkBinary(arr, target)
	binaryTime := time.Since(start)

	fmt.Printf("Busca Linear: índice %d em %v\n", linearIndex, linearTime)
	fmt.Printf("Busca Binária: índice %d em %v\n", binaryIndex, binaryTime)
	fmt.Printf("Busca binária foi %.2fx mais rápida!\n",
		float64(linearTime.Nanoseconds())/float64(binaryTime.Nanoseconds()))
}

func stepByStepDemo() {
	list := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	fmt.Printf("Buscando %d no array: %v\n", target, list)
	checkBinaryWithSteps(list, target)
}

func checkBinaryWithSteps(list []int, target int) int {
	low := 0              // Índice inicial (limite inferior)
	high := len(list) - 1 // Índice final (limite superior)
	step := 1             // Contador de passos

	for low <= high {
		mid := (low + high) / 2

		fmt.Printf("Passo %d: low=%d, high=%d, mid=%d, valor=%d\n",
			step, low, high, mid, list[mid])

		if list[mid] == target {
			fmt.Printf("🎉 Encontrado no índice %d!\n", mid)
			return mid
		}

		if list[mid] < target {
			// Elemento está na metade direita
			fmt.Printf("   %d < %d, buscando na metade direita\n", list[mid], target)
			low = mid + 1 // Descarta metade esquerda
		} else {
			// Elemento está na metade esquerda
			fmt.Printf("   %d > %d, buscando na metade esquerda\n", list[mid], target)
			high = mid - 1 // Descarta metade direita
		}
		step++
	}

	fmt.Printf("❌ Não encontrado\n")
	return -1
}

// Busca linear para comparação de performance
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// Busca binária com loop
func checkBinary(list []int, i int) int {
	low := 0              // Ponteiro para o início do intervalo
	high := len(list) - 1 // Ponteiro para o fim do intervalo

	for low <= high {
		mid := (low + high) / 2 // Calcula o índice do meio

		// Verifica se encontrou o elemento
		if list[mid] == i {
			return mid // Retorna o índice
		}

		// Decide qual metade manter
		if list[mid] < i {
			low = mid + 1 // Busca na metade direita
		} else {
			high = mid - 1 // Busca na metade esquerda
		}
	}
	return -1
}

// Busca binária com recursão
func recursiveCheckBinary(list []int, item int, high, low int) int {
	if high >= low {
		mid := (low + high) / 2 // Calcula o índice do meio

		// Caso base: encontrou o elemento
		if list[mid] == item {
			return mid
		} else if list[mid] > item {
			// Elemento está na metade esquerda
			// Chama recursivamente com novo limite superior
			return recursiveCheckBinary(list, item, mid-1, low)
		} else {
			// Elemento está na metade direita
			// Chama recursivamente com novo limite inferior
			return recursiveCheckBinary(list, item, high, mid+1)
		}
	}
	return -1
}
