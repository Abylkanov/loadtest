package main

import (
	"fmt"
	"github.com/alem-platform/ap"
)

func importS04() {
	ap.PutRune('a')
	fmt.Println(`test`)
}

type Stack struct {
	array []interface{}
  }
  
  func NewStack() *Stack {
	
	return &Stack{array:[]interface{}{}}
  }
  
  func (s *Stack) Push(v interface{}) {
	s.array=append(s.array, v)
  }
  
  func (s *Stack) Pop() interface{} {
	if len(s.array) == 0 {
		return nil
	}
	last:=len(s.array)-1
	item :=s.array[last]
	s.array = s.array[:last]
	return item
  }


  type Lock struct {
    status bool
}

func (l *Lock) Lock() bool {
	if !l.status {
		l.status = true
		return true
	}
	return false
}

func (l *Lock) Unlock() bool {
	if l.status {
		l.status = false
		return true
	}
	return false

}

func (l *Lock) IsLocked() bool {
	return l.status

}

func NewLock() *Lock {
	return &Lock{status: false}
}

type Set struct {

	cart []string

}

func NewSet() *Set {
	return &Set{cart:[]string{}}
}

func (s *Set) Add(v string) {
	s.cart=append(s.cart, v)
}

func (s *Set) Has(v string) bool {
	for _, pos:=range s.cart {
		if pos == v {
			return true
		}
	}
	return false
}

func (s *Set) Delete(v string) {

	for i:=0; i < len(s.cart); i++ {
		if v == s.cart[i] {
			s.cart = append(s.cart[:i], s.cart[i+1:]...)
		}
	}

}

func (s *Set) Values() []string {
	return s.cart

}

func (s *Set) Difference(other *Set) *Set {
	var res Set
	hasClone:=false
	for _, item1:=range s.cart {
		for _,item2:=range other.cart {
			if item1==item2 {
				hasClone = true
			}
		}
		if hasClone == false {
			res.cart=append(res.cart,item1)
		}
	}
	return &res
}

func (s *Set) Intersection(other *Set) *Set {
	var res Set

	hasClone:=false
	for _, item1:=range s.cart {
		for _,item2:=range other.cart {
			if item1==item2 {
				hasClone = true
			}
		}
		if hasClone{
			res.cart=append(res.cart,item1)
		}
	}
	return &res
	
}

func CoinsCombinations(sum int) [][]int {
    coins := []int{100, 50, 10, 5, 1}
    
    var result [][]int
    
    var findCombinations func(target int, current []int, start int)
    findCombinations = func(target int, current []int, start int) {
        if target == 0 {
            combination := make([]int, len(current))
            copy(combination, current)
            result = append(result, combination)
            return
        }
        
        for i := start; i < len(coins); i++ {
            coin := coins[i]
            
            if coin <= target {
                current = append(current, coin)
                
                findCombinations(target-coin, current, i)
                
                current = current[:len(current)-1]
            }
        }
    }
    
    findCombinations(sum, []int{}, 0)
    
    return result
}

func isUnique(s string) bool {
	chars := make(map[rune]bool)
	for _, c := range s {
		if chars[c] {
			return false
		}
		chars[c] = true
	}
	return true
}

func UniqCombN(s string, n int) []string {
	// Если строка имеет дубликаты или невозможно сформировать комбинации, возвращаем пустой срез
	if !isUnique(s) || n > len(s) {
		return []string{}
	}

	// Преобразуем строку в срез рун (для корректной работы с Unicode)
	runes := []rune(s)
	var result []string

	// Вложенная функция для генерации комбинаций
	var generateCombinations func(current []rune, start int)
	generateCombinations = func(current []rune, start int) {
		if len(current) == n {
			// Когда длина текущей комбинации равна n, добавляем её в результат
			result = append(result, string(current))
			return
		}

		// Перебираем возможные символы для добавления
		for i := start; i < len(runes); i++ {
			// Добавляем символ в текущую комбинацию
			current = append(current, runes[i])
			// Рекурсивный вызов для следующего символа
			generateCombinations(current, i+1)
			// Убираем последний символ для возвращения на предыдущий уровень
			current = current[:len(current)-1]
		}
	}

	// Генерируем все комбинации длины n
	generateCombinations([]rune{}, 0)

	return result
}

type ListNode struct {
	Next *ListNode
	Value interface{}
  }
  
  type List struct {
	Head, Tail *ListNode
  }
  
  func NewList() *List {
	return &List{}
  
  }
  
  func (l *List) PushBack(v interface{}) {
	newNode:= &ListNode{Value:v}

	if l.Tail == nil {
		l.Head=newNode
		l.Tail=newNode
	} else {
		l.Tail.Next = newNode
		l.Tail=newNode
	}

  
  }
  
  func (l *List) Reverse() {
    var prev *ListNode
    current := l.Head
    var next *ListNode
    
    // Traverse the list and reverse the direction of each node
    for current != nil {
        next = current.Next      // Store the next node
        current.Next = prev      // Reverse the link
        prev = current           // Move prev and current one step forward
        current = next
    }
    
    // After the loop, prev will be the new Head of the list
    l.Head, l.Tail = l.Tail, l.Head
}

  
func (l *List) Sort(fn func(a *ListNode, b *ListNode) int) {
    if l.Head == nil {
        return // Если список пуст, ничего не делаем
    }

    // Простая сортировка вставками для списка
    for current := l.Head; current != nil; current = current.Next {
        for next := current.Next; next != nil; next = next.Next {
            // Если элементы в неправильном порядке, меняем их местами
            if fn(current, next) > 0 {
                current.Value, next.Value = next.Value, current.Value
            }
        }
    }
}

// ForEach выполняет действие для каждого элемента списка
func (l *List) ForEach(fn func(n *ListNode)) {
    for current := l.Head; current != nil; current = current.Next {
        fn(current)
    }
}

type BTree struct {
	Root *BTreeNode
  }
  
  type BTreeNode struct {
	Parent *BTreeNode
	Left, Right *BTreeNode
	Value int
  }
  
  func NewBTree() *BTree {
	return &BTree{}
  }
  
  func (b *BTree) ReplaceOrInsert(v int) {
	if b.Root == nil {
		// Если дерево пустое, создаем корень
		b.Root = &BTreeNode{Value: v}
		return
	}

	// Ищем подходящее место для вставки
	current := b.Root
	for {
		if v < current.Value {
			// Если значение меньше текущего, идем в левое поддерево
			if current.Left == nil {
				current.Left = &BTreeNode{Value: v, Parent: current}
				break
			}
			current = current.Left
		} else if v > current.Value {
			// Если значение больше текущего, идем в правое поддерево
			if current.Right == nil {
				current.Right = &BTreeNode{Value: v, Parent: current}
				break
			}
			current = current.Right
		} else {
			// Если значение уже существует, заменяем его
			current.Value = v
			break
		}
	}
}

// ApplyByLevel выполняет функцию fn для каждого узла на каждом уровне дерева
func (b *BTree) ApplyByLevel(fn func(node *BTreeNode, level int)) {
	if b.Root == nil {
		return
	}

	// Используем очередь для обхода дерева по уровням
	queue := []*BTreeNode{b.Root}
	level := 0

	for len(queue) > 0 {
		// Обрабатываем текущий узел
		currentNode := queue[0]
		queue = queue[1:]

		// Выполняем функцию для текущего узла
		fn(currentNode, level)

		// Добавляем детей в очередь
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}

		// Если очередь пуста, переходим к следующему уровню
		if len(queue) > 0 && queue[0].Parent != currentNode.Parent {
			level++
		}
	}
}

func getLast_(s []string) string {
	if len(s) == 0 {
		return ""
	} else {
		return s[len(s)-1]
	}
}

func toUpper_GameCities(s byte) rune {
	if rune(s) >= 'a' && rune(s) <= 'z' {
		return rune(s - 32)
	} else {
		return rune(s)
	}
}

func solve_GameCities(ans *[]string, used map[string]bool, cities []string, temp []string) {
	if len(temp) > len(*ans) {
		*ans = make([]string, len(temp))
		copy(*ans, temp)
	}

	for i := 0; i < len(cities); i++ {
		if getLast_(temp) == "" {
			used[cities[i]] = true
			temp = append(temp, cities[i])
			solve_GameCities(ans, used, cities, temp)
			used[cities[i]] = false
			temp = temp[:len(temp)-1]
		} else {
			s := getLast_(temp)
			if toUpper_GameCities(s[len(s)-1]) == toUpper_GameCities(cities[i][0]) && !used[cities[i]] {
				used[cities[i]] = true
				temp = append(temp, cities[i])
				solve_GameCities(ans, used, cities, temp)
				used[cities[i]] = false
				temp = temp[:len(temp)-1]
			}
		}
	}
}

func GameCities(cities []string) []string {
	used := map[string]bool{}
	ans := []string{}
	solve_GameCities(&ans, used, cities, []string{})
	return ans
}
