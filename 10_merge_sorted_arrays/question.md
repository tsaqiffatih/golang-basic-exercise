### Merge Sorted Arrays
**Description:**

Buat fungsi bernama MergeSortedArrays yang menerima dua array yang sudah diurutkan, dan menggabungkannya menjadi satu array yang juga sudah diurutkan.

**Example Input and Output:**

```go
MergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6}) // Output: []int{1, 2, 3, 4, 5, 6}
MergeSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}) // Output: []int{1, 2, 3, 4, 5, 6}
MergeSortedArrays([]int{5, 6, 7}, []int{1, 2, 3}) // Output: []int{1, 2, 3, 5, 6, 7}
MergeSortedArrays([]int{}, []int{1, 2, 3})        // Output: []int{1, 2, 3}
```