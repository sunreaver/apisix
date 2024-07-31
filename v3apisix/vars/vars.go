package vars

// 由一个或多个 [var, operator, val] 元素组成的列表。
// 例如：
// [["a", "=", "1"], ["b", "=", "2"]]
// 表示变量 a 的值是 1，变量 b 的值是 2。
type Vars [][]string
