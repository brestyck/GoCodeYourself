def build_heap(array, n, i):
    largest = i
    left_child = 2 * i + 1
    right_child = 2 * i + 2

    if left_child < n and array[i] < array[left_child]:
        largest = left_child

    if right_child < n and array[largest] < array[right_child]:
        largest = right_child

    if largest != i:
        (array[i], array[largest]) = (array[largest], array[i])  # swap
        build_heap(array, n, largest)
    print(array)


def heapSort(array):
    length = len(array)

    for element in range(length // 2 - 1, -1, -1):
        build_heap(array, length, element)

    for element in range(length - 1, 0, -1):
        array[element], array[0] = array[0], array[element]
        build_heap(array, element, 0)


arr = [12, 11, 13, 5, 6, 7]
heapSort(arr)
print("Sorted array is")
print(*arr)
